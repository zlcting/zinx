package ziface

//IRouter 路由抽象接口
//路由里的数据都是irequest
type IRouter interface {
	//再处理conn业务之前的钩子方法hook
	PreHandle(request IRequest)
	//处理conn业务的住方法hook
	Handle(request IRequest)
	//处理con业务之后的钩子方法hook
	PostHandle(request IRequest)
}
