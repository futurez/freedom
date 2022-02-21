package finterface

type CondResult int32

const (
	Cond_InConform = CondResult(0) //不符合
	Cond_Conform   = CondResult(1) //符合
	Cond_LastOne   = CondResult(2) //最后一个
)

type IConnManager interface {
	Add(conn IConnection)                                                    //添加链接
	Delete(connId int64)                                                     //删除链接
	Clear()                                                                  //删除并停止所有链接
	Get(connId int64) (IConnection, error)                                   //获取链接
	GetCond(cond func(conn IConnection) CondResult) ([]IConnection, error)   //根据条件获取conn
	Range(f func(connId int64, conn IConnection))                            //遍历链接
	Len() int                                                                //获取链接长度
	Broadcast(msg IMessage)                                                  //广播
	BroadcastCond(msg IMessage, f func(connId int64, conn IConnection) bool) //有条件广播, f返回true广播
}
