package ziface

import "context"

type IChain interface {
	Proceed(interface{}) interface{}
	Handler
	Middleware
}
type Handler func(ctx context.Context, req IRequest)
type Middleware func(Handler) Handler
