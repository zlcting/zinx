package znet

import "net"
import "zinx/zinxServer/ziface"

//Connection 链接模块
type Connection struct {
	//当前链接的socket tcp 套接字
	Conn *net.TCPConn
	//链接ID
	ConnID uint32
	//当前链接状态
	isClosed bool
	//当前链接所绑定的处理业务的方法API
	handleAPI ziface.HandFunc

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool
}

//NewConnection 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI ziface.HandFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

func (c *Connection) start() {

}

func (c *Connection) Stop() {}

//GetTCPConnection 获取当前链接的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//GetConnID 获取当前链接id
func (c *Connection) GetConnID() uint32 {
	return c.GetConnID()
}

//RemoteAddr 获取远程客户端的状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//Send 发送数据 将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {

}
