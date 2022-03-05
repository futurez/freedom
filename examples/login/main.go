package main

import (
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/fmessage"
	"github.com/futurez/freedom/fnet"
	"github.com/futurez/freedom/futils"
)

// GWsSvr 服务器唯一链接
var GWsSvr finterface.IServer

// GetSvrConnListByType 通过类型获取服务器链接列表,(一个类型服务器有多个实力,如游戏服务器,接入服务器等)
func GetSvrConnListByType(svrType int32) []finterface.IConnection {
	list, _ := GWsSvr.GetConnManager().GetCond(func(conn finterface.IConnection) finterface.CondResult {
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

// GetSvrConnByType 通过服务器类型获取当个服务器,（单体服务器）
func GetSvrConnByType(svrType int32) finterface.IConnection {
	list, err := GWsSvr.GetConnManager().GetCond(func(conn finterface.IConnection) finterface.CondResult {
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
	flog.InitLogger("login", eproto.LOGIN_SERVER, "./log", flog.DebugLevel)

	defer flog.SyncLogger()
	// 监听端口
	GWsSvr = fnet.NewWsServer("/msg", fmessage.NewJsonPack(), &ServerConn{})

	InitRouter()

	// 主动链接管理服务器
	svrClient := NewServerClient(eproto.MANAGER_SERVER, 1, "127.0.0.1", 9999)
	wsClient := fnet.NewWsClient("127.0.0.1", 9999, "/msg", true, svrClient, GWsSvr)
	wsClient.ConnectWebSocket()

	futils.Graceful()
}
