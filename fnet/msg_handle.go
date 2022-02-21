package fnet

import (
	"github.com/futurez/freedom/finterface"
)

type MsgHandle struct {
}

func (b MsgHandle) PreHook(ctx finterface.IContext) {
	//flog.Debugf("[MsgHandle] start [%d]\n", ctx.GetMsgId())
}

func (b MsgHandle) Handle(ctx finterface.IContext) {
	//flog.Debugf("[MsgHandle] base handle [%d]\n", ctx.GetMsgId())
}

func (b MsgHandle) PostHook(ctx finterface.IContext) {
	//flog.Debugf("[MsgHandle] end [%d]\n", ctx.GetMsgId())
}

type HandlerFunc func(ctx finterface.IContext)

func (b HandlerFunc) PreHook(ctx finterface.IContext) {
	//flog.Debugf("[HandlerFunc] start [%d]\n", ctx.GetMsgId())
}

func (b HandlerFunc) Handle(ctx finterface.IContext) {
	//flog.Debugf("[HandlerFunc] base handle [%d]\n", ctx.GetMsgId())
	b(ctx)
}

func (b HandlerFunc) PostHook(ctx finterface.IContext) {
	//flog.Debugf("[HandlerFunc] end [%d]\n", ctx.GetMsgId())
}
