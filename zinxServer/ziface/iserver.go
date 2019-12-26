package ziface

//IServer is interface
type IServer interface {
	Start()
	Stop()
	Serve()

	//路由功能，给当前的服务注册一个路由方法，供给客户端的链接处理使用
	AddRouter(msgID uint32, router IRouter)
}
