package ziface

//IMessage 定义一个抽象模块
type IMessage interface {
	//获取消息id
	GetMsgId() uint32
	//获取消息的长度
	GetMsgLen() uint32
	//获去消息内容
	GetData() []byte
	//设置消息id
	SetMsgId(uint32)
	//设置消息内容
	SetDate([]byte)
	//设置消息长度
	SetDataLen(uint32)
}
