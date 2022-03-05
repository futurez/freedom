package main

import (
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
)

// ServerConn 服务器链接
type ServerConn struct{}

// OnConnect 调用OnConnect
func (s *ServerConn) OnConnect(conn finterface.IConnection) {
}

func (s *ServerConn) OnDisconnect(conn finterface.IConnection) {
	if val, ok := conn.GetCache(eproto.CACHE_SVR_TYPE); !ok {
		flog.Warnf("not found svr type %d", val)
	} else {
		svrInfo := val.(eproto.ServerInfo)
		// 通知其他相关服务器，当前服务器端口
		flog.Debugf("disconnect svr_type:%d id:%d ip:%s:%d",
			svrInfo.ServerType, svrInfo.ServerId, svrInfo.ListenIp, svrInfo.ListenPort)
	}
	flog.Debug("Disconnect ")
}
