package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// PingRouter MsgId=1
type PingRouter struct {
	znet.BaseRouter
}

// Handle  Ping Handle MsgId=1
func (r *PingRouter) Handle(request ziface.IRequest) {
	//read client data
	// 也就是说这里的拆包规则要改成读取每一行数据 然后读取第三个占位符 然后再进行获取每一行有用的数据
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	// GetMsgID() 更改获取消息ID的方法
	// GetData() 更改获取消息的方法
}

// ziface主要是存放一些Zinx框架的全部模块的抽象层接口类，Zinx框架的最基本的是服务类接口iserver，定义在ziface模块中。

// 1.启动服务器的方法

// nginx 转发tcp请求

func main() {
	//1 Create a server service
	s := znet.NewServer()

	//2 configure routing
	s.AddRouter(1, &PingRouter{}) //

	//3 start service
	s.Serve()
}

// todo
// 1.系统退出的时候，需要关闭所有的链接

// 2.先接受到tcp XT的包

// 3.如何nginx只转发XT的请求到go上面

// 4.日志分割

// 每一行第三个标识位是tcp包类型
// 根据标识位去获取特定handler
// 先读取一行的数据来搞看看
// tcp的鉴权 日志

// 如何管理代码仓库 gitlab搭建
