package main

import (
	"encoding/json"
	"github.com/futurez/freedom/examples/eproto"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
)

func ServerRegister(context finterface.IContext) {
	msg := context.GetMessage()
	req := eproto.ServerRegisterReq{}
	if err := json.Unmarshal(msg.GetMsgData(), &req); err != nil {
		flog.Errorf("[ServerRegister] msg unmarshal err : %s", err.Error())
		return
	}
	// 缓存服务器信息
	context.GetConnection().SetCache(eproto.CACHE_SVR_TYPE, req.Info)

	flog.Infof("[ServerRegister] msgData:%+v", req.Info)
	// todo
	// 通知关联服务器,新增链接

	//返回注册通知
	resp := eproto.ServerRegisterResp{}
	// todo
	// 返回关联服务器信息

	context.GetConnection().SendMsgJson(eproto.CMD_SERVER_REGISTER_RESP, 0, resp)
}

func InitRouter() {
	// 添加消息处理
	G_WsServer.AddHandleFunc(eproto.CMD_SERVER_REGISTER_REQ, ServerRegister)
}
