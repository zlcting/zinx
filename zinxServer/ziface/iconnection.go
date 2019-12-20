package ziface

import "net"

// IConnection 定义连接接口
type IConnection interface {
	Start()
	Stop()
	//获取当前链接的socket conn
	GetTCPConnection() *net.TCPConn
	//获取当前链接id
	GetConnID() uint32

	//获取远程客户端的状态
	RemoteAddr() net.Addr

	//发送数据 将数据发送给远程的客户端
	Send(data []byte) error
}

// HandFunc 定义一个统一处理链接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
