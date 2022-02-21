package finterface

type IContext interface {
	GetConnection() IConnection
	// 获取链接ID
	GetConnId() int64
	// 消息
	GetMessage() IMessage
	// 业务消息ID
	GetMsgId() uint32
	// 业务消息Buffer
	GetMsgData() []byte
	// 业务消息长度
	GetMsgLen() uint32
	// 错误码
	GetErrCode() int32
}
