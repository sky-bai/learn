package ziface

// IRouter 路由配置抽象类
// 现在我们来给Zinx实现一个非常简单基础的路由功能，目的当然就是为了快速的让Zinx步入到路由的阶段。后续我们会不断的完善路由功能。
// 路由就是真正处理的实际逻辑

type IRouter interface {
	// PreHandle 在处理conn业务之前的钩子方法Hook
	PreHandle(request IRequest)
	// Handle 在处理conn业务的主方法Hook
	Handle(request IRequest)
	// PostHandle 在处理conn业务之后的钩子方法Hook
	PostHandle(request IRequest)
}

// 我们知道router实际上的作用就是，服务端应用可以给Zinx框架配置当前链接的处理业务方法，之前的Zinx-V0.2我们的Zinx框架处理链接请求的方法是固定的，现在是可以自定义，并且有3种接口可以重写。
// 当然每个方法都有一个唯一的形参IRequest对象，也就是客户端请求过来的连接和请求数据，作为我们业务方法的输入数据。
