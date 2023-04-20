package main

import "learn/55_zinx/zinx/znet"

func main() {
	s := znet.Server{Name: "zinx v0.1", IPVersion: "tcp4", IP: ""}
	s.Serve()

}

import (
	"learn/55_zinx/zinx/znet"
)

/*
模拟客户端
*/
func main() {

	/*
		服务端测试
	*/
	//1 创建一个server 句柄 s
	s := znet.NewServer()

	s.AddRouter(0, &znet.PingRouter{})
	//s.AddRouter(1, &znet.HelloZinxRouter{})
	/*
		客户端测试
	*/

	//2 开启服务
	s.Serve()
}
