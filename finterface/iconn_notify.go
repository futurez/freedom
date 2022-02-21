package finterface

// 链接对象通知
type IConnNotify interface {
	//调用OnConnect
	OnConnect(conn IConnection)
	//调用OnDisconnect
	OnDisconnect(conn IConnection)
}
