package finterface

type IServer interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//获取消息解析器
	GetMsgPack() IMsgPack
	//获取路由
	GetRouter() IRouter
	// 添加消息路由接口
	AddHandle(msgId uint32, handle IMsgHandle)
	// 添加消息路由方法
	AddHandleFunc(msgId uint32, handle func(IContext))
	//链接管理
	GetConnManager() IConnManager
}
