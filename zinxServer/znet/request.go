package znet

import "zinx/zinxServer/ziface"

//Request a
type Request struct {
	//已经和客户端建立好链接的conn
	conn ziface.IConnection
	//客户端请求的数据
	//data []byte
	msg ziface.IMessage
}

//GetConnection 得到当前链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//GetData 得到当前数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//GetMsgID a
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
