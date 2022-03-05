package finterface

type IMsgHandle interface {
	PreHandle(ctx IContext)  //处理业务之前钩子
	Handle(ctx IContext)     //处理conn业务的方法
	PostHandle(ctx IContext) //处理之后钩子
}
