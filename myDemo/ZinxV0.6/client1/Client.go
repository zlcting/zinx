package main

import "fmt"

import "time"

import "net"
import "zinx/zinxServer/znet"

import "io"

func main() {
	fmt.Println("client start")
	//1 直接链接远程服务器 得到一个链接
	time.Sleep(1 * time.Second)
	//2 链接调用write 写数据
	conn, err := net.Dial("tcp", "127.0.0.1:8999")

	if err != nil {
		fmt.Println(err)
	}

	for {
		//发送封包的消息格式
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(1, []byte("zinxv client Test Message")))
		if err != nil {
			fmt.Println("pack err", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error", err)
			return
		}
		//服务器该回复一个message 数据，msgId：1 pingpingping

		//先读取流中的head的部分 得到id 和datalen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			break
		}
		//将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msghead error", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			//2根据datalen进行第二次读取，将data读出来
			msg := msgHead.(*znet.Message)

			msg.Data = make([]byte, msg.DataLen)
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("READ msg data error,", err)
				return
			}
			fmt.Println("-->Recv Server Msg :ID = ", msg.Id, ",len = ", msg.DataLen, ",data = ", string(msg.Data))

		}

		//在根据datalen进行第二次读取，将data读出来

		time.Sleep(1 * time.Second)

	}
}
