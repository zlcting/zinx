package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/zinxServer/utils"
	"zinx/zinxServer/ziface"
)

//Connection 链接模块
type Connection struct {
	//当前conn 隶属于哪个server
	TCPServer ziface.IServer

	//当前链接的socket tcp 套接字
	Conn *net.TCPConn
	//链接ID
	ConnID uint32
	//当前链接状态
	isClosed bool
	//当前链接所绑定的处理业务的方法API
	// handleAPI ziface.HandFunc

	//该链接处理的方法
	//Router ziface.IRouter

	//消息的管理msgid 对应的业务api
	MsgHandler ziface.IMsgHandler

	//告知当前链接已经退出的/停止 channel (有reader 告知writer退出)
	ExitChan chan bool
	//无缓冲的管道   用于读写协成的通信
	msgChan chan []byte
}

//NewConnection 初始化链接模块的方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		TCPServer: server,
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		//handleAPI: callbackAPI,
		MsgHandler: msgHandler,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
	}

	//将conn加入到connmananger
	c.TCPServer.GetConnMr().Add(c)
	return c
}

//StartReader 启动从当前链接读数据的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutione is runing")
	defer fmt.Println("connId =", c.ConnID, "[reader is exit],remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		// buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("recv buf err", err)
		// 	continue
		// }
		//创建拆包 解包的对象
		dp := NewDataPack()
		//读取客户端的msg head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		//拆包 得到msgID 和 msgDatalen 放再msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("Unpack error", err)
			break
		}

		// 根据datalen 再次读取data 放再msg.data 中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}

		msg.SetDate(data)

		//得到当前conn数据的request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已开启了工作池 将消息发送给woker工作池处理即可
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//调用执行注册的路由方法
			//从路由中，找到注册绑定的conn对应的router
			go c.MsgHandler.DoMsgHandler(&req)
		}

	}
}

//StartWriter 开启回写的协成 写消息
func (c *Connection) StartWriter() {
	fmt.Println("[writer groutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn writer is exit]")

	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("send data error,", err)
				return
			}
		case <-c.ExitChan:
			//代表reader已经退出 此时writer也要退出
			return

		}
	}
}

//Start as
func (c *Connection) Start() {
	fmt.Println("conn start() connID = ", c.ConnID)

	//启动当前链接的读数据业务
	//todo 启动从当前链接写数据的业务
	go c.StartReader()
	//todo 启动从当前链接写数据业务
	go c.StartWriter()

	//按照用户传递进来的创建连接时需要处理的业务,执行钩子方法
	c.TCPServer.CallOnConnStart(c)
	//==================
}

//Stop as
func (c *Connection) Stop() {
	fmt.Println("conn stop connid = ", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//如果用户注册了该链接的关闭回调业务,那么在此刻应该显示调用
	c.TCPServer.CallOnConnStop(c)
	//关闭socket链接
	c.Conn.Close()

	//告知writer 关闭
	c.ExitChan <- true

	//将当前连接从connmr中摘除
	c.TCPServer.GetConnMr().Remove(c)
	//回收资源
	close(c.ExitChan)
	close(c.msgChan)

}

//GetTCPConnection 获取当前链接的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//GetConnID 获取当前链接id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//RemoteAddr 获取远程客户端的状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//SendMsg 发送数据 将数据先封包再发送给远程的客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection close where send msg")
	}

	//将data 进行封包 msgdatalen msgid msgdata

	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("pack error msg id = ", msgId)
		return errors.New("pack error msg")
	}
	//将数据写回客户端
	c.msgChan <- binaryMsg

	return nil
}
