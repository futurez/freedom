package finterface

const (
	CONNECTING = int32(1) //正在链接
	CONNECTED  = int32(2) //链接成功
	CLOSING    = int32(3) //正在断开链接
	CLOSED     = int32(4) //已经断开链接
)

type IConnection interface {
	Connected()                                                   //成功链接
	Close() bool                                                  //关闭链接 (ture)
	GetConnID() int64                                             //获取链接ID
	RemoteAddr() string                                           //远程地址
	SendMessage(msg IMessage) error                               //发送消息
	SendMsgData(msgId uint32, code int32, data []byte) error      //发送消息
	SendMsgJson(msgId uint32, code int32, data interface{}) error //发送json数据
	GetConnStats() int32                                          //获取链接状态

	SetCache(key string, val interface{})    //设置缓存
	GetCache(key string) (interface{}, bool) //获取缓存
}
