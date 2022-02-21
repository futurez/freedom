package main

import (
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
)

// 服务器链接
type ServerClient struct {
	SvrInfo eproto.ServerInfo
}

//调用OnConnect
func (s *ServerClient) OnConnect(conn finterface.IConnection) {
	conn.SetCache(eproto.CACHE_SVR_TYPE, s.SvrInfo)

	//str := fmt.Sprintf("serverType=%d", s.ServerType)
	//
	//conn.SendMsgData(1, 0, []byte(str))
	//flog.Debug("Connect Send ", str)
}

func (s *ServerClient) OnDisconnect(conn finterface.IConnection) {
	if val, ok := conn.GetCache(eproto.CACHE_SVR_TYPE); !ok {
		flog.Warnf("not found svr type %d", val)
	} else {
		svrInfo := val.(eproto.ServerInfo)
		// 主动链接对方的网络断开，处理相关业务
		flog.Debugf("disconnect svr_type:%d id:%d ip:%s:%d",
			svrInfo.ServerType, svrInfo.ServerId, svrInfo.ListenIp, svrInfo.ListenPort)
	}
	flog.Debug("Disconnect ")
}
