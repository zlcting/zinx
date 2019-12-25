package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx/zinxServer/utils"
	"zinx/zinxServer/ziface"
)

//DataPack 封包 拆包的具体模块
type DataPack struct {
}

//NewDataPack 初始化
func NewDataPack() *DataPack {
	return &DataPack{}
}

//GetHeadLen 获取包的头长度
func (dp *DataPack) GetHeadLen() uint32 {
	//datalen uint32 （4字节）+ID uint32（4个字节）
	return 8
}

//Pack 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放byte字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	//将dataLen写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//将MsgId 写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err

	}
	//将data数据 写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//Unpack 拆包方法 将包的head信息读出来 之后再根据信息里的data的长度 再进行一次读
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioreader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head信息，得到datalen和msgId
	msg := &Message{}

	//读datalen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//读msgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断datalen 是否已经超出了我们允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		fmt.Println(utils.GlobalObject.MaxPacketSize)
		return nil, errors.New("too large msg data recv")
	}

	return msg, nil

}
