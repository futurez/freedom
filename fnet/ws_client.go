package fnet

import (
	"fmt"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"github.com/gorilla/websocket"
	"net/url"
	"time"
)

type WsClient struct {
	addr        string                 //(ws://ip:port/path)
	ip          string                 //监听IP
	port        int32                  //监听端口
	path        string                 //请求url (ws://ip:port/path)
	IsReconnect bool                   //是否重新链接
	notify      finterface.IConnNotify //链接状态变更通知
	server      finterface.IServer     //添加到的服务器

	finterface.IConnection // ws connect
}

func NewWsClient(ip string, port int32, path string, bReconnect bool, notify finterface.IConnNotify, server finterface.IServer) *WsClient {
	c := WsClient{
		addr:        fmt.Sprintf("%s:%d", ip, port),
		ip:          ip,
		port:        port,
		path:        path,
		IsReconnect: bReconnect,
		notify:      notify,
		server:      server,
		IConnection: nil,
	}
	if c.server == nil {
		c.server = DefWsServer
	}
	return &c
}

func (c *WsClient) ConnectWebSocket() {
	u := url.URL{Scheme: "ws", Host: c.addr, Path: c.path}

	go func() {
		for {
			flog.Debugf("[freedom] connecting to %s", u.String())
			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				flog.Warnf("[freedom] connecting to %s err:%s", u.String(), err.Error())
				if !c.IsReconnect {
					flog.Infof("[freedom] : not reconnect, quit remote %s", u.String())
					break
				}
				time.Sleep(30 * time.Second)
				continue
			} else {
				connId := generateConnId()
				c.IConnection = newWsConnection(c.server, conn, connId, c)
				//c.WsConn = wsConn.(*WsConn)
				c.IConnection.Connected()
				break
			}
		}
	}()
}

func (c *WsClient) OnConnect(conn finterface.IConnection) {
	flog.Debugf("[freedom] WsClient %s connected", conn.RemoteAddr())
	if c.notify != nil {
		c.notify.OnConnect(c)
	}
}

//调用OnDisconnect
func (c *WsClient) OnDisconnect(conn finterface.IConnection) {
	flog.Debugf("WsClient %s disconnect", conn.RemoteAddr())
	if c.notify != nil {
		c.notify.OnDisconnect(c)
	}

	if c.IsReconnect {
		c.ConnectWebSocket()
	}
}

func (c *WsClient) Close() bool {
	flog.Debugf("WsClient %s close.", c.RemoteAddr())
	if c.IConnection != nil {
		c.IConnection.Close()
	}
	c.IsReconnect = false
	return true
}
