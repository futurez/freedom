package main

import (
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/fconf"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
)

// ServerClient 服务器链接
type ServerClient struct {
	eproto.ServerInfo
}

func NewServerClient(svrType, svrId int32, ip string, port int32) *ServerClient {
	return &ServerClient{eproto.ServerInfo{
		ServerType: svrType,
		ServerId:   svrId,
		ListenIp:   ip,
		ListenPort: port,
	}}
}

func (s *ServerClient) OnConnect(conn finterface.IConnection) {
	conn.SetCache(eproto.CACHE_SVR_TYPE, s.ServerInfo)
	req := eproto.ServerRegisterReq{
		Info: eproto.ServerInfo{
			ServerType: eproto.LOGIN_SERVER,
			ServerId:   fconf.Conf.ServerId,
			ListenIp:   fconf.Conf.WebsocketIP,
			ListenPort: fconf.Conf.WebsocketPort,
		},
	}
	conn.SendMsgJson(eproto.CMD_SERVER_REGISTER_REQ, 0, req)
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
