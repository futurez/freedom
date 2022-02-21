package fnet

import (
	"github.com/futurez/freedom/finterface"
)

type BaseClient struct {
}

func (c *BaseClient) OnConnect(conn finterface.IConnection) {
	//	flog.Debugf("BaseClient %s connected", conn.RemoteAddr())
}

//调用OnDisconnect
func (c *BaseClient) OnDisconnect(conn finterface.IConnection) {
	//flog.Debugf("BaseClient %s disconnect", conn.RemoteAddr())
}
