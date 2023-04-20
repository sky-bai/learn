package main

import (
	"fmt"
	"learn/92_zinx/zinx/ziface"
	"learn/92_zinx/zinx/znet"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	//创建一个server句柄
	s := znet.NewServer("[zinx V0.3]")

	s.AddRouter(&PingRouter{})

	//2 开启服务
	s.Serve()
}
