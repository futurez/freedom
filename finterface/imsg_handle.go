package finterface

type IMsgHandle interface {
	PreHook(ctx IContext)		//处理业务之前钩子
	Handle(ctx IContext)		//处理conn业务的方法
	PostHook(ctx IContext)		//处理之后钩子
}