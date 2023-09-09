package znet

import (
	"backend/bin/tcp/ziface"
	"context"
	"time"
)

//// Handler defines the handler invoked by Middleware.
//type Handler func(ctx context.Context, req interface{})
//
//// 定义中间件要处理的handler类型
//
//// 实际业务具体的handler处理函数
//
//// Middleware is HTTP/gRPC transport middleware.
//type Middleware func(Handler) Handler

// 中间件类型处理这个函数

// 然后要思考每一层的处理逻辑

// 对具体函数的中间件函数处理方法 这一层也是函数处理 这里只需要做一层抽象 具体的中间件操作通过高阶函数去处理
// 链式调用 可以 和 装饰器模式来组合使用

// Chain returns a Middleware that specifies the chained handler for endpoint.
func Chain(m ...ziface.Middleware) ziface.Middleware { // 返回中间件的实际处理逻辑
	return func(next ziface.Handler) ziface.Handler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next) // 最后一个函数处理过后的handler 传入 下一个 函数     相当于是接受最后一个中间件的处理 然后将这个处理的结果传入到前一个中间件的入参中
			// 将实际的handler传入到中间件中 并将处理后的结果赋值给前一个中间件的入参
		}
		return next // 最后得到的next就是整个中间件链的入口，也就是最终的处理器。 然后一个被所有中间件处理过后的handler
	}
}

func TimeoutMiddleware(second uint) ziface.Middleware {
	return func(handler ziface.Handler) ziface.Handler {
		return func(ctx context.Context, req ziface.IRequest) {
			t := time.Duration(second) * time.Second
			ctx, cancel := context.WithTimeout(ctx, t)
			defer cancel()
			handler(ctx, req)
			return
		}
	}
}
