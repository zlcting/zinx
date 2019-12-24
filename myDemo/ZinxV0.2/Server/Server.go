package main

import "zinx/zinxServer/znet"

func main() {
	s := znet.NewServer("[zinx v0.2]")
	s.Serve()
}
