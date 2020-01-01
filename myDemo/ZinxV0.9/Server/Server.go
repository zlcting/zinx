package main

import (
	"fmt"
	"zinx/zinxServer/ziface"
	"zinx/zinxServer/znet"
)

//PingRouter ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Handle test handle
func (t *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call PingRouter router Handle")
	//先读取客户端的数据 再回写ping。。。。ping。。。。ping
	fmt.Println("recv from client :msgid=", request.GetMsgID(), "data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte("ping.....ping....ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

//Handle test handle
func (t *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("call HelloRouter router Handle")
	//先读取客户端的数据 再回写ping。。。。ping。。。。ping
	fmt.Println("recv from client :msgid=", request.GetMsgID(), "data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("hello welcome to zinxTCP"))
	if err != nil {
		fmt.Println(err)
	}
}

//DoConnectionBegin 创建链接后的hook
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("===>doconnectionBegin is called")
	if err := conn.SendMsg(202, []byte("do msg begin")); err != nil {
		fmt.Println(err)
	}

}

//DoConnectionLost a
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("=======>do connection lost")
	fmt.Println("conn id = ", conn.GetConnID(), "is lost")
}

func main() {
	s := znet.NewServer("[zinx v0.5]")

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	//添加路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
