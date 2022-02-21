package eproto

const (
	MANAGER_SERVER = 1 //管理服务器
	LOGIN_SERVER   = 2 //登录服务器
)

type ServerInfo struct {
	ServerType int32  `json:"t"`
	ServerId   int32  `json:"d"`
	ListenIp   string `json:"i"`
	ListenPort int32  `json:"p"`
}

type ServerRegisterReq struct {
	Info ServerInfo `json:"i"`
}

type ServerRegisterResp struct {
	Connected []ServerInfo `json:"c"` //需要链接的对象
}
