package fconf

import (
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/futils"
)

type Config struct {
	ServerName     string `json:"serverName"`     //服务器名称
	WebsocketIP    string `json:"websocketIP"`    //websocketIP
	WebsocketPort  int32  `json:"websocketPort"`  //websocket端口
	ServerType     int32  `json:"serverType"`     //服务器类型(1-1024)
	ServerId       int32  `json:"serverId"`       //服务器ID(1-32)
	MaxWsConn      int    `json:"maxWsConn"`      //websocket最大链接数
	WorkerPoolSize uint32 `json:"workerPoolSize"` //业务工作池的数量
	WorkerMsgCap   uint32 `json:"workerMsgSize"`  //每一个工作池最大消息队列
	SendMsgChanLen uint32 `json:"sendMsgChanLen"` //每一个链接发送消息缓存队列
}

var Conf Config

func init() {
	if err := futils.ReadConfig("config.json", &Conf); err != nil {
		flog.Warnf("[freedom] ReadConfig app.cfg err:%s and use default config", err.Error())
		Conf = Config{
			WebsocketIP:    "0.0.0.0",
			WebsocketPort:  8080,
			ServerType:     1,
			ServerId:       1,
			MaxWsConn:      100000,
			WorkerPoolSize: 10,
			WorkerMsgCap:   1000,
			SendMsgChanLen: 50,
		}
	} else {
		if Conf.MaxWsConn <= 100 {
			Conf.MaxWsConn = 100000
		}
	}
	//flog.Debugf("%+v", Conf)
}
