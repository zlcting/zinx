package znet

import "zinx/zinxServer/ziface"

import "sync"

import "fmt"

import "errors"

type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理的连接集合
	connLock    sync.RWMutex                  //保护连接结合的读写锁
}

//NewConnManager 创建当前连接的方法
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//Add 添加链接
func (connMr *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源map 加写锁
	connMr.connLock.Lock()
	defer connMr.connLock.Unlock()

	connMr.connections[conn.GetConnID()] = conn
	fmt.Println("connid = ", conn.GetConnID(), "connection add to connmanager syccessfully conn num = ", connMr.Len())
}

//Remove 删除链接
func (connMr *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源map 加写锁
	connMr.connLock.Lock()
	defer connMr.connLock.Unlock()
	//删除map连接信息
	delete(connMr.connections, conn.GetConnID())

	fmt.Println("connid = ", conn.GetConnID(), "connection Remove to connmanager syccessfully conn num = ", connMr.Len())
}

//Get 根据connid获取链接
func (connMr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//保护共享资源map 读写锁
	connMr.connLock.RLock()
	defer connMr.connLock.RUnlock()

	if conn, ok := connMr.connections[connID]; ok {

		return conn, nil
	}

	return nil, errors.New("connction not found")
}

//Len 得到当前链接总数
func (connMr *ConnManager) Len() int {
	return len(connMr.connections)
}

//ClearConn 清楚并终止所有的链接
func (connMr *ConnManager) ClearConn() {
	//保护共享资源map 加写锁
	connMr.connLock.Lock()
	defer connMr.connLock.Unlock()

	//删除conn 并停止conn的工作
	for connID, conn := range connMr.connections {
		conn.Stop()

		delete(connMr.connections, connID)
	}
	fmt.Println("clear all connections succ conn num = ", connMr.Len())
}
