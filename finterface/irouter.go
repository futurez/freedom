package finterface

type IRouter interface {
	// 发送网络消息
	SendConnMsg(ctx IConnection, msg IMessage)
	// 开启工作池
	StartWorkerPool()
	// 关闭工作池
	StopWorkerPool()
	// 添加消息路由接口
	AddHandle(msgId uint32, handle IMsgHandle)
	// 添加消息路由方法
	AddHandleFunc(msgId uint32, handle func(IContext))
}
