package ziface

//消息管理抽象层

//IMsgHandler 调度执行对应的路由 消息处理方法
type IMsgHandler interface {
	//为消息添加具体的处理逻辑
	DoMsgHandler(request IRequest)
	//为消息添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)

	StartWorkerPool()
	//将消息发送给消息人物队列
	SendMsgToTaskQueue(request IRequest)
}
