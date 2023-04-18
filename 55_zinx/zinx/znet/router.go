package znet

import "learn/55_zinx/zinx/ziface"

// BaseRouter 实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct{}

// 这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化

func (br *BaseRouter) PreHandle(req ziface.IRequest)  {}
func (br *BaseRouter) Handle(req ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(req ziface.IRequest) {}

//  IServer增添路由添加功能
//
// 我们需要给IServer类，增加一个抽象方法AddRouter,目的也是让Zinx框架使用者，可以自定一个Router处理业务方法。

//  我们之前在已经给Zinx配置了路由模式，但是很惨，之前的Zinx好像只能绑定一个路由的处理业务方法。显然这是无法满足基本的服务器需求的，那么现在我们要在之前的基础上，给Zinx添加多路由的方式。
//
// 既然是多路由的模式，我们这里就需要给MsgId和对应的处理逻辑进行捆绑。所以我们需要一个Map。

// 将消息类型与对应的路由进行绑定

// 这里起名字是Apis，其中key就是msgId， value就是对应的Router，里面应是使用者重写的Handle等方法。
//
// 那么这个Apis应该放在哪呢。
//
// 我们再定义一个消息管理模块来进行维护这个Apis。
