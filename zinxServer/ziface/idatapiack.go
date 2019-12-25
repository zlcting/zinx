package ziface

//IDataPack 封包 拆包 模块
type IDataPack interface {
	//获取包的长度
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	Unpack([]byte) (IMessage, error)
}
