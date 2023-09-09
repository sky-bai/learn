package znet

import (
	"backend/bin/tcp/ziface"
	"backend/lib/logger"
	"context"
	"errors"
	"sync"
)

var _ ziface.IConnManager = (*ConnManager)(nil)

/*
ConnManager 连接管理模块
*/
type ConnManager struct {
	connections map[string]ziface.IConnection //管理的连接信息
	connLock    sync.RWMutex                  //读写连接的读写锁
}

/*
NewConnManager 创建一个链接管理
*/
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[string]ziface.IConnection),
	}
}

// Add 添加链接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//将conn连接添加到ConnManager中
	connMgr.connections[conn.GetConnID()] = conn

	logger.Infof(context.Background(), "connection add to ConnManager successfully: conn num = %d", connMgr.Len())
}

// Remove 删除连接
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除连接信息
	delete(connMgr.connections, conn.GetConnID())

	logger.Infof(context.Background(), "connection Remove ConnID=%s,successfully: conn num = %d", conn.GetConnID(), connMgr.Len())
}

// Get 利用ConnID获取链接
func (connMgr *ConnManager) Get(connID string) (ziface.IConnection, error) {
	//保护共享资源Map 加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// Len 获取当前连接
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

// ClearConn 清除并停止所有连接
func (connMgr *ConnManager) ClearConn() {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range connMgr.connections {
		//停止
		conn.Stop()
		//删除
		delete(connMgr.connections, connID)
	}

	logger.Infof(context.Background(), "Clear All Connections successfully: conn num = %d", connMgr.Len())
}
