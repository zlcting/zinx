package ziface

//IServer is interface
type IServer interface {
	Start()
	Stop()
	Serve()

	//路由功能，给当前的服务注册一个路由方法，供给客户端的链接处理使用
	AddRouter(msgID uint32, router IRouter)
	GetConnMr() IConnManager
	SetOnConnStart(func(IConnection))
	//设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(IConnection))
	//调用连接OnConnStart Hook函数
	CallOnConnStart(conn IConnection)
	//调用连接OnConnStop Hook函数
	CallOnConnStop(conn IConnection)
}
