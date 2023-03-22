package kratos

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	err     error
	wg      sync.WaitGroup
	errOnce sync.Once

	workerOnce sync.Once
	ch         chan func(ctx context.Context) error // 1.定义一个channel 保存需要执行的函数
	chs        []func(ctx context.Context) error

	ctx    context.Context
	cancel func()
}

// WithContext create a Group.
// given function from Go will receive this context,
func WithContext(ctx context.Context) *Group {
	return &Group{ctx: ctx}
}

// WithCancel create a new Group and an associated Context derived from ctx.
//
// given function from Go will receive context derived from this ctx,
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithCancel(ctx context.Context) *Group {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{ctx: ctx, cancel: cancel}
}

func (g *Group) do(f func(ctx context.Context) error) { // 每一个do方法接受一个具有上下文的函数
	ctx := g.ctx // 1.获取本身的ctx，没有的话就自己设置一个
	if ctx == nil {
		ctx = context.Background()
	}
	var err error
	defer func() { // 2.defer 先不看 是函数执行完之后的操作
		if r := recover(); r != nil { // 3.如果有panic，就把panic的信息转换成error
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}
		if err != nil {
			g.errOnce.Do(func() { // 只获取一次错误
				g.err = err
				if g.cancel != nil { // 如果有cancel，就取消
					g.cancel()
				}
			})
		}
		g.wg.Done()
	}()
	err = f(ctx) // 3.调用执行函数
}

// GOMAXPROCS set max goroutine to work.
func (g *Group) GOMAXPROCS(n int) {
	if n <= 0 {
		panic("errgroup: GOMAXPROCS must great than 0")
	}
	g.workerOnce.Do(func() {
		g.ch = make(chan func(context.Context) error, n)
		for i := 0; i < n; i++ {
			go func() {
				for f := range g.ch {
					g.do(f)
				}
			}()
		}
	})
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func(ctx context.Context) error) {
	g.wg.Add(1)
	if g.ch != nil {
		select {
		case g.ch <- f:
		default:
			g.chs = append(g.chs, f)
		}
		return
	}
	go g.do(f)
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	if g.ch != nil {
		for _, f := range g.chs {
			g.ch <- f
		}
	}
	g.wg.Wait()
	if g.ch != nil {
		close(g.ch) // let all receiver exit 如果都执行完了就关闭channel
	}
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

// wait 方法会阻塞直到所有的函数调用都返回，然后返回第一个非空的错误
// 1.先去判断channel是否为空，如果不为空，就把数组厘米的元素都放入channel中
