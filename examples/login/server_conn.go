package main

import (
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
)

// 被动接受的服务器链接
type ServerConn struct {
	SvrInfo eproto.ServerInfo
}

//调用OnConnect
func (s *ServerConn) OnConnect(conn finterface.IConnection) {
}

//被动链接断开
func (s *ServerConn) OnDisconnect(conn finterface.IConnection) {
	if val, ok := conn.GetCache(eproto.CACHE_SVR_TYPE); !ok {
		flog.Warnf("not found svr type %d", val)
	} else {
		svrInfo := val.(eproto.ServerInfo)
		// 某个服务器链接断开,处理相关业务
		flog.Debugf("disconnect svr_type:%d id:%d ip:%s:%d",
			svrInfo.ServerType, svrInfo.ServerId, svrInfo.ListenIp, svrInfo.ListenPort)
	}
	flog.Debug("Disconnect ")
}
