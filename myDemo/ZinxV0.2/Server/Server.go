package main

import zent "zinx/zinxServer/znet"

func main() {
	s := zent.NewServer("[zinx v0.2]")
	s.Serve()
}
