package fnet

import (
	"errors"
	"fmt"
	"github.com/futurez/freedom/finterface"
	"github.com/futurez/freedom/flog"
	"sync"
)

type ConnManager struct {
	mapConnections map[int64]finterface.IConnection
	lock           sync.RWMutex
}

func newConnManager() finterface.IConnManager {
	return &ConnManager{mapConnections: make(map[int64]finterface.IConnection)}
}

//添加链接
func (m *ConnManager) Add(conn finterface.IConnection) {
	m.lock.Lock()
	defer m.lock.Unlock()

	connId := conn.GetConnID()
	m.mapConnections[connId] = conn

	flog.Infof("[freedom] <add> %d | %d ", connId, len(m.mapConnections))
}

//删除链接
func (m *ConnManager) Delete(connId int64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.mapConnections, connId)

	flog.Debug("connection delete ConnID=", connId, " successfully: conn num = ", len(m.mapConnections))
}

//删除并停止所有链接
func (m *ConnManager) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	for k, conn := range m.mapConnections {
		//停止
		conn.Close()
		delete(m.mapConnections, k)
	}
	flog.Debugf("Clear All connections successfully: conn num = %d", len(m.mapConnections))
}

//获取链接
func (m *ConnManager) Get(connId int64) (finterface.IConnection, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if conn, ok := m.mapConnections[connId]; ok {
		return conn, nil
	}
	return nil, errors.New(fmt.Sprintf("not found connId=%d", connId))
}

//根据条件获取conn
func (m *ConnManager) GetCond(cond func(finterface.IConnection) finterface.CondResult) (list []finterface.IConnection, err error) {
	if cond == nil {
		return nil, errors.New(fmt.Sprintf("cond is nil"))
	}
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, v := range m.mapConnections {
		if res := cond(v); res == finterface.Cond_Conform {
			list = append(list, v)
		} else if res == finterface.Cond_LastOne {
			list = append(list, v)
			break
		}
	}
	if len(list) > 0 {
		return list, nil
	}
	return nil, errors.New(fmt.Sprintf("not through the cond"))
}

//遍历链接
func (m *ConnManager) Range(f func(connId int64, conn finterface.IConnection)) {
	if f == nil {
		return
	}
	m.lock.RLock()
	defer m.lock.RUnlock()
	for k, v := range m.mapConnections {
		f(k, v)
	}
}

//获取链接长度
func (m *ConnManager) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.mapConnections)
}

//广播
func (m *ConnManager) Broadcast(msg finterface.IMessage) {
	m.Range(func(connId int64, conn finterface.IConnection) {
		conn.SendMessage(msg)
	})
}

//有条件广播, f返回true广播
func (m *ConnManager) BroadcastCond(msg finterface.IMessage, cond func(connId int64, conn finterface.IConnection) bool) {
	m.Range(func(connId int64, conn finterface.IConnection) {
		if cond != nil {
			if !cond(connId, conn) { // false 不发送
				return
			}
		}
		conn.SendMessage(msg)
	})
}
