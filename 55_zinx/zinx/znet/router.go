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
