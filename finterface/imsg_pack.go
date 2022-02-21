package finterface

type PackType uint32

const (
	PACK_UNKNOWN = PackType(0)
	PACK_BINARY  = PackType(1) // msg  : head+ body 字节流 (用于tcp)
	PACK_JSON    = PackType(2) // json : {"id": 1, "d" : []byte } (用于websocket)
	PACK_PROTO   = PackType(3) // eproto: 						  (用于websocket)
)

type IMsgPack interface {
	GetPackType() PackType
	GetHeadLen() uint32 //TCP Message Head
	Pack(IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
