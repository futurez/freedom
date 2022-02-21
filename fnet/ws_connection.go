package fnet

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/futurez/freedom/fconf"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/fmessage"
	"github.com/gorilla/websocket"
	"sync"
	"sync/atomic"
)

type WsConn struct {
	// 当前Conn属于哪个Server
	ws finterface.IServer
	// 当前websocket链接对象
	conn *websocket.Conn
	// 当前链接ID,也可以称作SessionId,ID全服唯一
	connId int64
	// 链接状态
	connStats int32
	// 处理读写协程上下文
	connCtx    context.Context
	connCancel context.CancelFunc
	// 有缓存区消息管道 用于read,write两个goroutine之间消息通信
	sendMsgChan chan finterface.IMessage
	//链接状态变更
	connNotify finterface.IConnNotify

	//链接缓存
	cache map[string]interface{}
	//保护当前cache的锁
	cacheLock sync.RWMutex
}

func newWsConnection(server finterface.IServer, wsSocket *websocket.Conn, connId int64, notify finterface.IConnNotify) finterface.IConnection {
	conn := &WsConn{
		ws:          server,
		conn:        wsSocket,
		connId:      connId,
		connStats:   finterface.CONNECTING,
		sendMsgChan: make(chan finterface.IMessage, fconf.GConf.SendMsgChanLen),
		cache:       make(map[string]interface{}),
		connNotify:  notify,
	}
	return conn
}

func (c *WsConn) GetConnStats() int32 {
	return atomic.LoadInt32(&c.connStats)
}

func (c *WsConn) GetConnID() int64 {
	return c.connId
}

func (c *WsConn) RemoteAddr() string {
	return c.remoteAddr()
}

func (c *WsConn) remoteAddr() string {
	if c.conn != nil {
		return c.conn.RemoteAddr().String()
	}
	return ""
}

func (c *WsConn) Connected() {
	if !atomic.CompareAndSwapInt32(&c.connStats, finterface.CONNECTING, finterface.CONNECTED) {
		return
	}
	//flog.Info(c.remoteAddr(), " Connected ConnId = ", c.connId, " stats = ", c.connStats)

	c.ws.GetConnManager().Add(c)

	// 处理读写协程上下文
	c.connCtx, c.connCancel = context.WithCancel(context.Background())

	var wait sync.WaitGroup

	// 先启动Write Goroutine, 保证链接成功可以发送消息
	wait.Add(1)
	go c.writePump(&wait)
	wait.Wait()

	// 通知应用层链接成功
	if c.connNotify != nil {
		c.connNotify.OnConnect(c)
	}

	// 启动Read Goroutine
	go c.readPump()
}

// close -> read end -> write end
// 主动关闭 (外部调用 OR readPump调用)
func (c *WsConn) Close() bool {
	if !atomic.CompareAndSwapInt32(&c.connStats, finterface.CONNECTED, finterface.CLOSING) {
		return false
	}

	flog.Info(c.remoteAddr(), " Close connId = ", c.connId, " stats = ", c.connStats)
	// 关闭readPump
	if c.connCancel != nil {
		c.connCancel()
	}
	// 关闭writePump
	if c.sendMsgChan != nil {
		close(c.sendMsgChan)
		c.sendMsgChan = nil
	}
	return true
}

// 被动关闭通知
func (c *WsConn) onClosed() bool {
	if !atomic.CompareAndSwapInt32(&c.connStats, finterface.CLOSING, finterface.CLOSED) {
		return false
	}

	// 调用链接断开
	if c.connNotify != nil {
		c.connNotify.OnDisconnect(c)
	}

	// 关闭链接
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}

	// 清空cache
	c.cache = nil

	// 从链接池删除
	c.ws.GetConnManager().Delete(c.connId)
	return true
}

func (c *WsConn) readPump() {
	defer func() {
		flog.Info(c.remoteAddr(), " [readPump] exit! connId:", c.connId)
		if !c.Close() {
			c.onClosed()
		}
	}()

	//	flog.Info(c.remoteAddr(), " [readPump] goroutine is running! connId = ", c.connId)

	for {
		select {
		case <-c.connCtx.Done():
			return

		default:
			{
				msgType, msgData, err := c.conn.ReadMessage()
				if err != nil {
					flog.Warnf("%s [readPump] wsConnId:%d ReadMessage err:%v", c.RemoteAddr(), c.connId, err)
					return
				}
				// 解析消息
				msg, err := c.ws.GetMsgPack().Unpack(msgData)
				if err != nil {
					flog.Errorf("%s [readPump] wsConnId:%d unpack msg=%s err=%s", c.RemoteAddr(), string(msgData), err.Error())
					continue
				}
				flog.Debugf("%s [readPump] wsConnId:%d type:%d len:%d", c.RemoteAddr(), c.connId, msgType, len(msgData))
				// 发送conn,msg
				c.ws.GetRouter().SendConnMsg(c, msg)
			}
		}
	}
}

func (c *WsConn) writePump(wait *sync.WaitGroup) {
	//flog.Info(c.remoteAddr(), " [writePump] goroutine is running! connId=", c.connId)

	wait.Done()
	for {
		select {
		//case <-c.connCtx.Done():
		//	flog.Debug(c.RemoteAddr().String(), " is close [context] connId=", c.connId)
		//	return
		case msg, ok := <-c.sendMsgChan:
			if !ok { //chan没有数据,且关闭ok才会返回false
				flog.Info(c.remoteAddr(), " [writePump] exit! check close sendMsgChan connId=", c.connId)
				// 数据全部发送完毕,关闭链接
				c.onClosed()
				return
			}

			data, err := c.ws.GetMsgPack().Pack(msg)
			if err != nil {
				flog.Errorf("%s [writePump] connId=%d Pack Message Err=%s", c.RemoteAddr(), c.connId, err.Error())
				goto ERROR
			}
			//websocket msg type 1:websocket.TextMessage, 2:websocket.BinaryMessage
			if err = c.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				flog.Errorf("%s [writePump] connId=%d WriteMessage Err=%s", c.RemoteAddr(), c.connId, err.Error())
				goto ERROR
			}
		}
	}

ERROR:
	if !c.Close() { // 主动关闭失败,调用被动关闭
		c.onClosed()
	}
}

func (c *WsConn) SendMessage(msg finterface.IMessage) (err error) {
	if !atomic.CompareAndSwapInt32(&c.connStats, finterface.CONNECTED, finterface.CONNECTED) {
		return errors.New("ERR_NOT_CONNECTED")
	}

	select {
	case c.sendMsgChan <- msg:
		// 发送成功
		flog.Debugf("Send MsgId:%d Code:%d Len:%d", msg.GetMsgId(), msg.GetErrCode(), msg.GetMsgLen())
	case <-c.connCtx.Done():
		err = errors.New("ERR_CONNECTION_LOSS")
	default:
		err = errors.New("ERR_SEND_MSG_CHAN_FULL")
	}
	return err
}

func (c *WsConn) SendMsgData(msgId uint32, code int32, data []byte) error {
	msg := fmessage.NewMessage(msgId, code, data)
	return c.SendMessage(msg)
}

func (c *WsConn) SendMsgJson(msgId uint32, code int32, msg interface{}) error {
	if data, err := json.Marshal(msg); err != nil {
		return err
	} else {
		return c.SendMsgData(msgId, code, data)
	}
}

func (c *WsConn) SetCache(key string, val interface{}) {
	c.cacheLock.Lock()
	defer c.cacheLock.Unlock()
	c.cache[key] = val
}

func (c *WsConn) GetCache(key string) (interface{}, bool) {
	c.cacheLock.RLock()
	defer c.cacheLock.RUnlock()
	val, ok := c.cache[key]
	return val, ok
}
