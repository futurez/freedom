package fmessage

import (
	"encoding/json"
	"errors"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
)

type jsonMsg struct {
	Id   uint32 `json:"i"` //msgId
	Code int32  `json:"c"` //错误码
	Body []byte `json:"b"` //消息数据
}

type JsonPack struct {
}

func NewJsonPack() finterface.IMsgPack {
	return &JsonPack{}
}

func (m *JsonPack) GetHeadLen() uint32 {
	return 0
}

func (m *JsonPack) GetPackType() finterface.PackType {
	return finterface.PACK_JSON
}

func (m *JsonPack) Pack(msg finterface.IMessage) ([]byte, error) {
	j := jsonMsg{
		Id:   msg.GetMsgId(),
		Code: msg.GetErrCode(),
		Body: msg.GetMsgData(),
	}
	if b, err := json.Marshal(j); err != nil {
		return nil, errors.New("json msg pack " + err.Error())
	} else {
		return b, nil
	}
}

func (m *JsonPack) Unpack(msgData []byte) (finterface.IMessage, error) {
	var j jsonMsg
	if err := json.Unmarshal(msgData, &j); err != nil {
		flog.Errorf("unpack json msg=%s err=%s", string(msgData), err.Error())
		return nil, errors.New("unpack json " + err.Error())
	}
	msg := NewMessage(j.Id, j.Code, j.Body)
	return msg, nil
}
