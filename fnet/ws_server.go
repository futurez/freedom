package fnet

import (
	"fmt"
	"github.com/futurez/freedom/fconf"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/fmessage"
	"github.com/gorilla/websocket"
	"net/http"
	"sync/atomic"
)

type WsServer struct {
	ip         string                  //监听IP
	port       int32                   //监听端口
	pattern    string                  //请求url (ws://ip:port/pattern)
	connMgr    finterface.IConnManager //链接对象管理
	msgPack    finterface.IMsgPack     //基础包(head+body)
	router     finterface.IRouter      //消息路由
	connNotify finterface.IConnNotify  //链接状态变更通知
	stats      int32                   //状态1:启动,0:终止
}

// NewWsServer msgpack : default eproto pack
func NewWsServer(ip string, port int32, pattern string, pack finterface.IMsgPack, notify finterface.IConnNotify) finterface.IServer {
	ws := WsServer{
		ip:      ip,
		port:    port,
		pattern: pattern,
		connMgr: newConnManager(),
		router:  newMsgRouter(),
	}
	if pack == nil {
		pack = fmessage.NewJsonPack()
	}
	ws.msgPack = pack
	if notify == nil {
		notify = &BaseClient{}
	}
	ws.connNotify = notify
	return &ws
}

func (ws *WsServer) serveWs(w http.ResponseWriter, r *http.Request) {
	var upgrade = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte("not websocket."))
		return
	}

	if ws.connMgr.Len() >= fconf.GConf.MaxWsConn {
		flog.Warnf("websocket connection is full, maxWsConn=%d", fconf.GConf.MaxWsConn)
		conn.Close()
		return
	}

	connId := generateConnId()
	wsConn := newWsConnection(ws, conn, connId, ws.connNotify)
	flog.Infof("websocket listen = ws://%s:%d%s new connId=%d remoteAddr=%s", ws.ip, ws.port, ws.pattern, connId, conn.RemoteAddr().String())
	wsConn.Connected()
}

func (ws *WsServer) Start() {
	if !atomic.CompareAndSwapInt32(&ws.stats, 0, 1) {
		flog.Errorf("WsServer already start...")
		return
	}
	flog.Infof("WsServer start...")

	ws.router.StartWorkerPool()
	go func() {
		http.HandleFunc(ws.pattern, func(w http.ResponseWriter, r *http.Request) {
			ws.serveWs(w, r)
		})
		addr := fmt.Sprintf("%s:%d", ws.ip, ws.port)
		flog.Info("Start listen ws://" + addr + ws.pattern)
		if err := http.ListenAndServe(addr, nil); err != nil {
			flog.Error("err:", err)
			return
		}
	}()
}

func (ws *WsServer) Stop() {
	if !atomic.CompareAndSwapInt32(&ws.stats, 1, 0) {
		flog.Errorf("WsServer already stop...")
		return
	}

	flog.Infof("stop ws://%s:%d:%s", ws.ip, ws.port, ws.pattern)

	ws.router.StopWorkerPool()
	ws.connMgr.Clear()
}

func (ws *WsServer) GetMsgPack() finterface.IMsgPack {
	return ws.msgPack
}

func (ws *WsServer) GetRouter() finterface.IRouter {
	return ws.router
}

// 添加消息路由接口
func (ws *WsServer) AddHandle(msgId uint32, handle finterface.IMsgHandle) {
	ws.router.AddHandle(msgId, handle)
}

// 添加消息路由方法
func (ws *WsServer) AddHandleFunc(msgId uint32, handle func(finterface.IContext)) {
	ws.router.AddHandleFunc(msgId, handle)
}

//链接管理
func (ws *WsServer) GetConnManager() finterface.IConnManager {
	return ws.connMgr
}

// import golang.org/x/net/websocket
//func (s *WsServer) newConnection(ws *websocket.Conn) {
//
//	//fmt.Println("new connection remoteAddr ", ws.RemoteAddr().String())
//	////fmt.Println("new connection localAddr ", ws.LocalAddr().String())
//	//var err error
//	//for {
//	//	var reply string
//	//	if err = websocket.Message.Receive(ws, &reply); err != nil {
//	//		fmt.Println("can't receive")
//	//		break
//	//	}
//	//	fmt.Println("receive back from login: " + reply)
//	//	msg := "received: " + reply
//	//	fmt.Println("sending to login: " + msg)
//	//	if err = websocket.Message.Send(ws, msg); err != nil {
//	//		fmt.Println("Can't send")
//	//		break
//	//	}
//	//}
//	//ws.Close()
//	//fmt.Println("close connection ", ws.RemoteAddr().String())
//}
//
//func (ws *WsServer) Start() {
//	//// 过滤跨域
//	//ws := websocket.Server{
//	//	Handler:   websocket.Handler(s.newConnection),
//	//	Handshake: func(config *websocket.Config, req *http.Request) error {
//	//		config.Origin, _ = websocket.Origin(config, req)
//	//		return nil
//	//	},
//	//}
//	//http.Handle(s.Pattern, ws)
//	go func () {
//		http.Handle(ws.Pattern, websocket.Handler(ws.newConnection))
//		addr := fmt.Sprintf("%s:%d", ws.Ip, ws.Port)
//		flog.Info("start listen ws://" + addr + ws.Pattern)
//		if err := http.ListenAndServe(addr, nil); err != nil {
//			flog.Error("err ", err)
//			return
//		}
//	}()
//}