package main

import (
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/fmessage"
	"github.com/futurez/freedom/fnet"
	"github.com/futurez/freedom/futils"
)

// 服务器唯一链接
var G_WsServer finterface.IServer

//  通过类型获取服务器链接
func GetSvrConnByType(svrType int32) finterface.IConnection {
	conn, err := G_WsServer.GetConnManager().GetCond(func(conn finterface.IConnection) bool {
		if val, ok := conn.GetCache(eproto.CACHE_SVR_TYPE); !ok {
			return false
		} else if val.(int32) == svrType {
			return true
		} else {
			return false
		}
	})
	if err != nil {
		return nil
	}
	return conn
}

func main() {
	flog.InitLogger("manager", eproto.MANAGER_SERVER, "./log", flog.DebugLevel)
	defer flog.SyncLogger()

	G_WsServer = fnet.NewWsServer("0.0.0.0", 9999, "/msg", fmessage.NewJsonPack(), &ServerConn{})
	G_WsServer.Start()
	// 添加消息路由
	InitRouter()

	futils.Graceful()
}
