package znet

import "zinx/zinxServer/ziface"

//BaseRouter 实现router先嵌入baserouter基类 根据需求对这个基类重写
type BaseRouter struct {
}

//这里之所以baseRouter的方法都为空 是因为有的router不希望有preHandle 和Posthandle这两个业务
//所以router全部继承baserouter 然后重写就可以了

//PreHandle 再处理conn业务之前的钩子方法hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

//Handle 处理conn业务的住方法hook
func (br *BaseRouter) Handle(request ziface.IRequest) {}

//PostHandle 处理con业务之后的钩子方法hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
