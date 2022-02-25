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

//  通过类型获取服务器链接列表,(一个类型服务器有多个实力,如游戏服务器,接入服务器等)
func GetSvrConnListByType(svrType int32) []finterface.IConnection {
	list, _ := G_WsServer.GetConnManager().GetCond(func(conn finterface.IConnection) finterface.CondResult {
		if val, ok := conn.GetCache(eproto.CACHE_SVR_TYPE); !ok {
			return finterface.Cond_InConform
		} else if val.(int32) == svrType {
			return finterface.Cond_Conform
		} else {
			return finterface.Cond_InConform
		}
	})
	return list
}

//  通过类型获取服务器链接
func GetSvrConnByType(svrType int32) finterface.IConnection {
	list, err := G_WsServer.GetConnManager().GetCond(func(conn finterface.IConnection) finterface.CondResult {
		if val, ok := conn.GetCache(eproto.CACHE_SVR_TYPE); !ok {
			return finterface.Cond_InConform
		} else if val.(int32) == svrType {
			return finterface.Cond_LastOne
		} else {
			return finterface.Cond_InConform
		}
	})
	if err != nil {
		return nil
	}
	return list[0]
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
