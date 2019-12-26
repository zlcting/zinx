package znet

import (
	"fmt"
	"strconv"
	"zinx/zinxServer/ziface"
)

//MsgHandle 消息处理模块的实现
type MsgHandle struct {
	//存放每个msgid 所对应的处理方法
	Apis map[uint32]ziface.IRouter
}

//NewMsgHandle 初始化/创建msghandle方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

//DoMsgHandler 为消息添加具体的处理逻辑
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//1.从request中找到msgid
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgID(), "is not found need add")
		return
	}
	//2.根据msgid调度对应的router的方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

//AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	//1 判断 当前msg绑定的API处理方法是否存在
	if _, ok := mh.Apis[msgID]; ok {
		//id已经存在
		panic("repeat api ,msgid = " + strconv.Itoa(int(msgID)))
	}
	//2 添加msg与api的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api msgId = ", msgID, "succ!")
}
