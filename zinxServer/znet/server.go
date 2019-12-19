package zent

import (
	"fmt"
	"net"
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
}

//Start is server start function
func (s *Server) Start() {
	fmt.Printf("[start] server listenner at IP :%s port:%d is starting\n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPversion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println(err)
			return
		}

		listenner, err := net.ListenTCP(s.IPversion, addr)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("start zinx server succ", s.Name, "succ listennting....")

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				continue
			}
			//已经与客户端链接
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println(err)
						continue
					}
					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println(err)
					}
					continue
				}
			}()
		}
	}()

}

//Stop is server Stop function
func (s *Server) Stop() {
	//todo 将一些服务器的资源、状态或者一些已经开辟的链接信息进行停止或者回收
}

//Serve is server res function
func (s *Server) Serve() {
	//启动服务器
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态
	select {}
}

//NewServer is server new obj function
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPversion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
