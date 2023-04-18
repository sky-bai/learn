package ziface

// IServer 定义一个服务器接口
type IServer interface {
	Start() // 启动服务器
	Stop()  // 停止服务器
	Serve() // 运行服务器
	// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId uint32, router IRouter)
}
