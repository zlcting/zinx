package main

import "fmt"

import "time"

import "net"

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
		_, err := conn.Write([]byte("hello Zinx"))
		if err != nil {
			fmt.Println(err)
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(" server call back : %s, cnt = %d\n", buf,
			cnt)

		time.Sleep(1 * time.Second)
	}
}
