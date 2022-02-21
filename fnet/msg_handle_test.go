package fnet

import (
	"freedom/finterface"
	"freedom/flog"
	"freedom/fmessage"
	"testing"
)

func HelloHandle(ctx finterface.IContext) {
	flog.Info("hello handle ")
}

func TestHandle(t *testing.T) {
	flog.InitLogger("fnet_handle_test", 0, "./log", flog.DebugLevel)
	defer flog.SyncLogger()

	msg := fmessage.NewMessage(1, 0, nil)
	ctx := allocateContext().(*Context)
	ctx.msg = msg
	
	hf := HandlerFunc(HelloHandle)
	hf.PreHook(ctx)
	hf.Handle(ctx)
	hf.PostHook(ctx)

	bf := MsgHandle{}
	bf.PreHook(ctx)
	bf.Handle(ctx)
	bf.PostHook(ctx)
}
