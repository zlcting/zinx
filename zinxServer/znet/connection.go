package znet

import "net"
import "zinx/zinxServer/ziface"
import "fmt"

//Connection 链接模块
type Connection struct {
	//当前链接的socket tcp 套接字
	Conn *net.TCPConn
	//链接ID
	ConnID uint32
	//当前链接状态
	isClosed bool
	//当前链接所绑定的处理业务的方法API
	// handleAPI ziface.HandFunc

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool

	//该链接处理的方法
	Router ziface.IRouter
}

//NewConnection 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//handleAPI: callbackAPI,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

//StartReader 启动从当前链接读数据的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutione is runing")
	defer fmt.Println("connId =", c.ConnID, "reader is exit,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		//调用当前链接绑定的handleApi
		// if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		// 	fmt.Println("ConnId", c.ConnID, "handle is err", err)
		// 	break
		// }

		//得到当前conn数据的request请求数据
		req := Request{
			conn: c,
			data: buf,
		}

		//调用执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		//从路由中，找到注册绑定的conn对应的router

	}
}

//Start as
func (c *Connection) Start() {
	fmt.Println("conn start() connID = ", c.ConnID)

	//启动当前链接的读数据业务
	//todo 启动从当前链接写数据的业务
	go c.StartReader()
	//todo 启动从当前链接写数据业务
}

//Stop as
func (c *Connection) Stop() {
	fmt.Println("conn stop connid = ", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//关闭socket链接
	c.Conn.Close()
	//回收资源
	close(c.ExitChan)

}

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
// func (c *Connection) Send(data []byte) {

// }
