package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/zinxServer/ziface"
)

/*
存储一切有关Zinx框架的全局参数,供其他模块使用
一些参数也可以通过 用户根据 zinx.json来配置
*/
type GlobalObj struct {
	TcpServer        ziface.IServer //当前Zinx的全局Server对象
	Host             string         //当前服务器主机IP
	TcpPort          int            //当前服务器主机监听端口号
	Name             string         //当前服务器名称
	Version          string         //当前Zinx版本号
	MaxPacketSize    uint32         //都需数据包的最大值
	MaxConn          int            //当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32         //当前业务工作的worker池子的groutine数量
	MaxWorkerTaskLen uint32         //允许用户最大队列任务
}

var GlobalObject *GlobalObj

//读取用户的配置文件
func (g *GlobalObj) Reload() {

	data, err := ioutil.ReadFile("../conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//将json数据解析到struct中
	//fmt.Printf("json :%s\n", data)
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
	提供init方法,默认加载
*/
func init() {
	//初始化GlobalObject变量,设置一些默认值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.5",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024, //每个worker对应的消息对了的任务的数量最大值
	}
	//从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
