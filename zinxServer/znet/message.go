package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

//NewMsgPackage 创建message 方法
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//GetMsgId 获取消息id
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

//GetMsgLen 获取消息的长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen

}

//GetData 获去消息内容
func (m *Message) GetData() []byte {
	return m.Data

}

//SetMsgId 设置消息id
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

//SetDate 设置消息内容
func (m *Message) SetDate(data []byte) {
	m.Data = data

}

//SetDataLen 设置消息长度
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
