package fmessage

import (
	"github.com/futurez/freedom/finterface"
)

type Message struct {
	msgId   uint32
	msgData []byte
	msgLen  uint32
	code    int32 //返回结果 0: success
}

func NewMessage(msgId uint32, code int32, msgData []byte) finterface.IMessage {
	return &Message{
		//msgType: msgType,
		msgId:   msgId,
		msgData: msgData,
		msgLen:  uint32(len(msgData)),
		code:    code,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.msgId
}

func (m *Message) GetMsgData() []byte {
	return m.msgData
}

func (m *Message) GetMsgLen() uint32 {
	return m.msgLen
}

func (m *Message) GetErrCode() int32 {
	return m.code
}
