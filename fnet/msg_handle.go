package fnet

import (
	"github.com/futurez/freedom/finterface"
)

type BaseHandle struct {
}

func (b *BaseHandle) PreHandle(ctx finterface.IContext) {
}

func (b *BaseHandle) Handle(ctx finterface.IContext) {

}

func (b *BaseHandle) PostHandle(ctx finterface.IContext) {
}

type HandlerFunc func(ctx finterface.IContext)

func (b HandlerFunc) PreHandle(ctx finterface.IContext) {
	//flog.Debugf("[HandlerFunc] start [%d]\n", ctx.GetMsgId())
}

func (b HandlerFunc) Handle(ctx finterface.IContext) {
	//flog.Debugf("[HandlerFunc] base handle [%d]\n", ctx.GetMsgId())
	b(ctx)
}

func (b HandlerFunc) PostHandle(ctx finterface.IContext) {
	//flog.Debugf("[HandlerFunc] end [%d]\n", ctx.GetMsgId())
}
