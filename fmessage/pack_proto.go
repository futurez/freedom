package fmessage

import (
	"errors"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/fmessage/pb3"
)

type ProtoPack struct{}

func NewProtoPack() finterface.IMsgPack {
	return &ProtoPack{}
}

func (m *ProtoPack) GetHeadLen() uint32 {
	return 0
}

func (m *ProtoPack) GetPackType() finterface.PackType {
	return finterface.PACK_PROTO
}

func (m *ProtoPack) Pack(msg finterface.IMessage) ([]byte, error) {
	p := pb3.ProtoMessage{
		MsgId: msg.GetMsgId(),
		Code:  msg.GetErrCode(),
		Body:  msg.GetMsgData(),
	}
	if b, err := p.Marshal(); err != nil {
		return nil, errors.New("eproto msg pack " + err.Error())
	} else {
		return b, nil
	}
}

func (m *ProtoPack) Unpack(msgData []byte) (finterface.IMessage, error) {
	msg := pb3.ProtoMessage{}
	if err := msg.Unmarshal(msgData); err != nil {
		flog.Errorf("unpack proto3 msg=%s err=%s", string(msgData), err.Error())
		return nil, errors.New("unpack proto3 " + err.Error())
	}
	return NewMessage(msg.MsgId, msg.Code, msg.Body), nil
}
