package znet

import (
	"fmt"
	"net"
	"zinx/zinxServer/utils"
	"zinx/zinxServer/ziface"
)

//Server 定义一个server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定ip版本
	IPversion string
	//服务器监听的IP
	IP string

	Port int

	//Router ziface.IRouter
	//当前server的消息管理模块 用来绑定msgid和对应的处理业务api关系
	MsgHandler ziface.IMsgHandler
	//该server的连接管理器
	ConnMr ziface.IConnManager

	//该Server的连接创建时Hook函数
	OnConnStart func(conn ziface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn ziface.IConnection)
}

//Start is server start function
func (s *Server) Start() {
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,MaxPacketSize:%d\n", utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPacketSize)

	fmt.Printf("[start] server listenner at IP :%s port:%d is starting\n", s.IP, s.Port)

	go func() {
		//开启消息队列及worker工作池
		s.MsgHandler.StartWorkerPool()

		//1.获取一个tcpadd
		addr, err := net.ResolveTCPAddr(s.IPversion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPversion, addr)
		if err != nil {
			fmt.Println("listen", s.IPversion, "err", err)
			return
		}

		fmt.Println("start zinx server succ", s.Name, "succ listennting....")
		var cid uint32
		cid = 0
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			//设置最大链接个数的判断，如果超过最大链接，那么则关闭此新的链接
			if s.ConnMr.Len() >= utils.GlobalObject.MaxConn {
				//todo 给客户端响应一个超出最大链接的错误包
				fmt.Println("too manty connections maxconn = ", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//将处理新的链接的业务方法和conn进行绑定 得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			//启动 当前的链接业务处理
			go dealConn.Start()
		}
	}()

}

//Stop is server Stop function
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
	//todo 将一些服务器的资源、状态或者一些已经开辟的链接信息进行停止或者回收
	fmt.Println("[stop server name]", s.Name)
	s.ConnMr.ClearConn()
}

//Serve is server res function
func (s *Server) Serve() {
	//启动服务器
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态
	select {}
}

//AddRouter s
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)

	fmt.Println("add router succ!!")

}

//NewServer is server new obj function
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPversion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		//Router:    nil,
		MsgHandler: NewMsgHandle(),
		ConnMr:     NewConnManager(),
	}

	return s
}

//GetConnMr 获取链接管理句柄
func (s *Server) GetConnMr() ziface.IConnManager {
	return s.ConnMr
}

//SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

//SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

//CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

//CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}
