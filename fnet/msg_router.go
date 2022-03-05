package fnet

import (
	"github.com/futurez/freedom/fconf"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/futils"
	"sync"
	"sync/atomic"
	"time"
)

// 消息处理路由器
type MsgRouter struct {
	mapHandle      map[uint32]finterface.IMsgHandle
	workerPoolSize int64                      //业务工作线程池数量
	workerMsgQueue []chan finterface.IContext //业务工作消息队列
	state          int32                      //1:start, 0:not start
	// context池
	pool sync.Pool
}

func newMsgRouter() finterface.IRouter {
	r := &MsgRouter{
		mapHandle:      make(map[uint32]finterface.IMsgHandle),
		workerPoolSize: int64(fconf.Conf.WorkerPoolSize),
		// 每一个worker对应一个queue
		workerMsgQueue: make([]chan finterface.IContext, fconf.Conf.WorkerPoolSize),
		state:          0,
	}
	r.pool.New = func() interface{} {
		return allocateContext()
	}
	return r
}

func (r *MsgRouter) AddHandle(msgId uint32, handle finterface.IMsgHandle) {
	r.mapHandle[msgId] = handle
}

func (r *MsgRouter) AddHandleFunc(msgId uint32, handle func(ctx finterface.IContext)) {
	r.mapHandle[msgId] = HandlerFunc(handle)
}

func (r *MsgRouter) DoMsgHandle(conn finterface.IConnection, msg finterface.IMessage) {
	flog.Debugf("[freedom] <doMsgHandler> <pool-get> %d | %d", conn.GetConnID(), msg.GetMsgId())
	ctx := r.pool.Get().(*Context)
	ctx.conn = conn
	ctx.msg = msg

	if r.workerPoolSize > 0 && atomic.CompareAndSwapInt32(&r.state, 1, 1) {
		workId := conn.GetConnID() % r.workerPoolSize
		r.workerMsgQueue[workId] <- ctx
	} else {
		// 这边不能用协程, 因为使用协程可能因为GPM导致客户端消息不是按照顺序处理
		r.doMsgHandler(ctx)
	}
}

func (r *MsgRouter) StartWorkerPool() {
	if r.workerPoolSize == 0 {
		return
	}
	if !atomic.CompareAndSwapInt32(&r.state, 0, 1) {
		flog.Warnf("Router already start worker pool")
		return
	}
	flog.Infof("[freedom] Router start worker pool size[%d]", r.workerPoolSize)

	for i := int64(0); i < r.workerPoolSize; i++ {
		r.workerMsgQueue[i] = make(chan finterface.IContext, fconf.Conf.WorkerMsgCap)
		go r.startWorkerFunc(i, r.workerMsgQueue[i])
	}
}

func (r *MsgRouter) StopWorkerPool() {
	if r.workerPoolSize == 0 {
		return
	}
	if !atomic.CompareAndSwapInt32(&r.state, 1, 0) {
		flog.Warnf("Router already stop worker pool")
		return
	}
	flog.Infof("Router stop worker pool")

	for i := int64(0); i < r.workerPoolSize; i++ {
		close(r.workerMsgQueue[i])
	}
}

func (r *MsgRouter) startWorkerFunc(workId int64, msgQueue <-chan finterface.IContext) {
	//flog.Infof("Worker ID = %d Start", workId)
	for {
		select {
		case ctx, ok := <-msgQueue:
			if !ok {
				flog.Infof("Worker ID = %d Stop chan msgQueue is close.", workId)
				return
			}
			if ctx != nil {
				r.doMsgHandler(ctx)
			}
		}
	}
}

func (r *MsgRouter) doMsgHandler(ctx finterface.IContext) {
	startTime := time.Now()
	defer func() {
		flog.Infof("[freedom] <doMsgHandler> <pool-put> %d | %d |%s", ctx.GetConnection().GetConnID(), ctx.GetMsgId(), time.Since(startTime).String())
		r.pool.Put(ctx)
		futils.RecoverFromPanic("doMsgHandler", nil)
	}()

	handler, ok := r.mapHandle[ctx.GetMsgId()]
	if !ok {
		flog.Errorf("handler msgId = %d is not found\n", ctx.GetMsgId())
		return
	}
	handler.PreHandle(ctx)
	handler.Handle(ctx)
	handler.PostHandle(ctx)
}
