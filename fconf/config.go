package fconf

type Config struct {
	ServerType     uint32 `json:"serverType"`     //服务器类型(1-1024)
	ServerId       uint32 `json:"serverId"`       //服务器ID(1-32)
	MaxWsConn      int    `json:"maxWsConn"`      //websocket最大链接数
	WorkerPoolSize uint32 `json:"workerPoolSize"` //业务工作池的数量
	WorkerMsgCap   uint32 `json:"workerMsgSize"`  //每一个工作池最大消息队列
	SendMsgChanLen uint32 `json:"sendMsgChanLen"` //每一个链接发送消息缓存队列
}

var GConf *Config

func init() {
	GConf = &Config{
		ServerType:     1,
		ServerId:       1,
		MaxWsConn:      10000,
		WorkerPoolSize: 10,
		WorkerMsgCap:   1000,
		SendMsgChanLen: 50,
	}
}
