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

//PreHandle test preRouter
func (t *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call router preHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping \n"))
	if err != nil {
		fmt.Println("call back before ping err")
	}
}

//Handle test handle
func (t *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping..........ping \n"))
	if err != nil {
		fmt.Println("call back  pingpingpingping err")
	}

}

//PostHandle test post
func (t *PingRouter) PostHandle(request ziface.IRequest) {

	fmt.Println("call router postHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping \n"))
	if err != nil {
		fmt.Println("call back after ping err")
	}

}

func main() {
	s := znet.NewServer("[zinx v0.3]")
	//添加路由
	s.AddRouter(&PingRouter{})
	s.Serve()
}
