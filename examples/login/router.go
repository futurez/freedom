package main

import (
	"encoding/json"
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
)

func ServerRegisterResp(context finterface.IContext) {
	msg := context.GetMessage()
	resp := eproto.ServerRegisterResp{}
	if err := json.Unmarshal(msg.GetMsgData(), &resp); err != nil {
		flog.Errorf("[ServerRegisterResp] msg unmarshal err : %s", err.Error())
		return
	}
	flog.Infof("[ServerRegisterResp] msgData:%+v", resp.Connected)
	// todo
	// 需要主动链接的服务器
}

func InitRouter() {
	GWsSvr.AddHandleFunc(eproto.CMD_SERVER_REGISTER_RESP, ServerRegisterResp)
}
