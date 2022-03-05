Freedom是一个基于Golang的并发服务器框架(网络使用的是websocket)

---
## freedom源码地址
### Github
    go get -u github.com/futurez/freedom

## 二、快速启动
```bash
# 克隆项目
$ git clone https://github.com/futurez/freedom.git

# 进入manager服务端样例目录
$ cd ./freedom/examples/manager

# 服务端编译
$ go build -o manager.bin

# 启动服务器
$ ./manager.bin

# 进入login服务器样例目录
$ cd ../login

# 服务器编译
$ go build -o login.bin

# 启动login服务器进行测试
$ ./login.bin
```

#### server
基于freedom框架开发的服务器应用，主函数步骤比较精简，最多只需要3步即可。
1. 创建server句柄
2. 配置自定义路由及业务
3. 启动服务

```go
package main

import (
    "github.com/futurez/freedom/fmessage"
    "github.com/futurez/freedom/fnet"
)

func main() {
    // 1. 创建一个server句柄
    s := fnet.NewWsServer("/msg", fmessage.NewJsonPack(), nil)

    // 2. 配置路由
    s.AddHandle(1, &PingHandle{})   // 路由消息对象
    s.AddHandleFunc(2, HelloHandleFunc) // 路由消息函数

    // 3. 开启服务
    s.Run()
}
```

其中自定义路由及业务配置方式如下：
```go
package main

import (
    "github.com/futurez/freedom/finterface"
    "github.com/futurez/freedom/flog"
    "github.com/futurez/freedom/fnet"
)

// PingHandle 消息处理对象
type PingHandle struct {
    fnet.MsgHandle
}

func (p *PingHandle) Handle(ctx finterface.IContext) {
    flog.Debugf("PingHandle MsgId[%d] - [%s]", ctx.GetMsgId(), string(ctx.GetMsgData()))

    ctx.GetConnection().SendMsgData(1, 0, []byte("Pong"))
}

// HelloHandleFunc 消息处理函数
func HelloHandleFunc(ctx finterface.IContext) {
    flog.Debugf("HelloHandleFunc MsgId[%d] - [%s]", ctx.GetMsgId(), string(ctx.GetMsgData()))

    ctx.GetConnection().SendMsgData(1, 0, []byte("Hello"))
}
```

#### client
freedom的消息处理采用，`[MsgLength]|[MsgID]|[Data]`的封包格式
```go
package main

import (
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/futurez/freedom/fnet"
	"time"
)

func main() {
	// 消息注册
	fnet.WsAddHandleFunc(1, func(ctx finterface.IContext) {
		flog.Debugf("Ping MsgId[%d] - [%s]", ctx.GetMsgId(), string(ctx.GetMsgData()))
		ctx.GetConnection().SendMsgData(1, 0, []byte("Pong"))
	})

	// 主动链接
	wc := fnet.NewWsClient("127.0.0.1", 8080, "/msg", true, &ServerClient{}, nil)
	wc.ConnectWebSocket()

	//
	time.Sleep(time.Minute)
}
```

客户端链接对象函数:
```go
package main

import "github.com/futurez/freedom/finterface"

type ServerClient struct {
}
func (s *ServerClient) OnConnect(conn finterface.IConnection) {
	conn.SendMsgData(1, 0, []byte("ping"))
	conn.SendMsgData(2, 0, []byte("hello"))
}
func (s *ServerClient) OnDisconnect(conn finterface.IConnection) {
}

```

### Freedom 配置文件
```json
{
  "serverName"    : "freedomServer",
  "websocketIP"   : "0.0.0.0",
  "websocketPort" : 8888,
  "serverType"    : 1,
  "serverId"      : 2,
  "maxWsConn"     : 10000,
  "workerPoolSize": 10,
  "workerMsgSize" : 100,
  "sendMsgChanLen": 100
}
```
`serverName`:服务器应用名称

`websocketIP`:服务器IP

`websocketPort`:服务器监听端口

`serverType`: 服务器类型

`serverId` : 服务器编号

`maxWsConn`:允许的客户端链接最大数量

`workerPoolSize`:工作任务池最大工作Goroutine数量

`workerMsgSize`: 工作任务池消息队列长度

`sendMsgChanLen`: 发送消息队列长度



### I.服务器模块Server

```go
    func DefaultWsServer() finterface.IServer

    func NewWsServer(pattern string, pack finterface.IMsgPack, notify finterface.IConnNotify) finterface.IServer
```
#### 创建一个freedom服务器句柄，该句柄作为当前服务器应用程序的主枢纽，包括如下功能：

#### 1)运行服务
```go
  func (s *WsServer) Run(addr ...string)()
```
#### 2)停止服务
```go
  func (s *WsServer) Stop()
```
#### 3)获取消息解析器
```go
  func (s *WsServer) GetMsgPack() finterface.IMsgPack
```
#### 4)获取路由
```go
  func (s *WsServer) GetRouter() finterface.IRouter
```
#### 5)添加消息路由对象
```go
  func (s *WsServer) AddHandle(msgId uint32, handle finterface.IMsgHandle)
```
#### 6)添加消息路由方法
```go
  func (s *WsServer) AddHandleFunc(msgId uint32, handle func(finterface.IContext))
```
#### 7)链接管理
```go
  func (s *WsServer)GetConnManager() finterface.IConnManager
```

### II.路由模块

#### 1) 路由基类
```go
  //实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
  type BaseHandle struct {}
  // 这里之所以 BaseHandle 的方法都为空，
  // 是因为有的Handle不希望有PreHandle或PostHandle
  // 所以Handle全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化
  func (b *MsgHandle) PreHandle(ctx finterface.IContext) {}
  func (b *MsgHandle) Handle(ctx finterface.IContext) {} 
  func (b *MsgHandle) PostHandle(ctx finterface.IContext) {}
```
#### 2) 路由函数
```go
  //实现type HandlerFunc func(ctx finterface.IContext)类型函数
  func MsgHandle(ctx finterface.IContext) {}
```

### III.链接模块
#### 1) 成功链接
```go
  func (c *WsConn) Connected()
```
#### 2) 关闭链接
```go
  func (c *WsConn) Close() bool
```
#### 3) 获取链接ID
```go 
  func (c *WsConn) GetConnID() int64
```
#### 4) 远程地址
```go
  func (c *WsConn) RemoteAddr() string
```
#### 5) 发送消息
```go
  func (c *WsConn) SendMessage(msg finterface.IMessage) (err error)
```
#### 6) 发送json数据
```go
  func (c *WsConn) SendMsgData(msgId uint32, code int32, data []byte) error 
```
#### 7) 获取链接状态
```go
  func (c *WsConn) GetConnStats()
```
#### 8)  设置缓存                                     
```go
  func (c *WsConn) SetCache(key string, val interface{})
```
#### 9) 
```go
  func (c *WsConn) GetCache(key string) (interface{}, bool)
```
