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
	fmt.Println("call router Handle")
	//先读取客户端的数据 再回写ping。。。。ping。。。。ping
	fmt.Println("recv from client :msgid=", request.GetMsgID(), "data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping.....ping....ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx v0.5]")
	//添加路由
	s.AddRouter(&PingRouter{})
	s.Serve()
}
