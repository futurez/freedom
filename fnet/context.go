package fnet

import "github.com/futurez/freedom/finterface"

type Context struct {
	conn finterface.IConnection
	msg  finterface.IMessage
}

func allocateContext() finterface.IContext {
	return &Context{}
}

func (c *Context) SetConnection(conn finterface.IConnection) {
	c.conn = conn
}

func (c *Context) SetMessage(msg finterface.IMessage) {
	c.msg = msg
}

func (c *Context) GetConnection() finterface.IConnection {
	return c.conn
}

func (c *Context) GetConnId() int64 {
	if c.conn != nil {
		return c.conn.GetConnID()
	}
	return 0
}

func (c *Context) GetMessage() finterface.IMessage {
	return c.msg
}

func (c *Context) GetMsgId() uint32 {
	return c.msg.GetMsgId()
}

func (c *Context) GetMsgData() []byte {
	return c.msg.GetMsgData()
}

func (c *Context) GetMsgLen() uint32 {
	return c.msg.GetMsgLen()
}

func (c *Context) GetErrCode() int32 {
	return c.msg.GetErrCode()
}
