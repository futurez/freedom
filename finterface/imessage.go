package finterface

type IMessage interface {
	// 业务消息ID
	GetMsgId()   uint32
	// 业务消息Buffer
	GetMsgData() []byte
	// 业务消息长度
	GetMsgLen() uint32
	// 错误码
	GetErrCode() int32
}

