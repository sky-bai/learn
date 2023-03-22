package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

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
func WithCancel(ctx context.Context) *Group {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{ctx: ctx, cancel: cancel}
}

// Group kratos 版本
type Group struct {
	err     error
	wg      sync.WaitGroup
	errOnce sync.Once

	workerOnce sync.Once
	ch         chan func(ctx context.Context) error
	chs        []func(ctx context.Context) error

	ctx    context.Context
	cancel func()
}

// kratos在基础版本的基础上添加了一个chan控制并发数量，一个slice来缓存为并发的函数指针。 缓存还未执行的函数

// GoMaxProc 设置并行数量
func (g *Group) GoMaxProc(n int) {
	if n <= 0 || n > runtime.NumCPU() {
		n = runtime.NumCPU()
	}
	g.workerOnce.Do(func() { // 这里不加会有什么影响吗 ？
		g.ch = make(chan func(context.Context) error, n)
		for i := 0; i < n; i++ {
			go func() {
				for f := range g.ch { // 开n个协程去读取chan中的函数指针，并执行函数逻辑 也就是说只有n个协程在读取管道中的函数指针并执行，这不就是消费者吗
					g.do(f)
				}
			}()
		}
	})
}

// Wait 等待任务执行完毕
func (g *Group) Wait() error {
	if g.ch != nil { // 判断所有的任务是否执行完成，如果并发chan里面没有就将缓存的函数指针放入chan中
		for _, f := range g.chs {
			g.ch <- f
		}
	}
	g.wg.Wait()
	if g.ch != nil {
		close(g.ch) // let all receiver exit
	}
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
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

type ABC struct {
	CBA int
}

func data() {
	var (
		abcs = make(map[int]*ABC)
		err  error
	)
	fmt.Println("start", time.Now())
	for i := 0; i < 10; i++ {
		abcs[i] = &ABC{CBA: i}
	}
	g := WithCancel(context.Background())
	g.Go(func(context.Context) (err error) {
		time.Sleep(time.Second * 5) // 实际任务耗时5s
		fmt.Println("task2")
		return nil
	})
	g.Go(func(context.Context) (err error) {
		//time.Sleep(time.Second * 1) // 实际任务耗时5s
		abcs[1].CBA++
		return errors.New("err1")
	})

	//time1 := time.AfterFunc(time.Second*1, func() {
	//	g.cancel()
	//	fmt.Println("timeout time:", time.Now())
	//})
	//defer time1.Stop()

	if err = g.Wait(); err != nil {
		fmt.Println("err", err)
		fmt.Println("end err:", time.Now())
		return
	}
	fmt.Println("end", time.Now())
}

func main() {
	//ReadTask()
	data()
	time.Sleep(20 * time.Second)
}
func WithContext(ctx context.Context) *Group {
	return &Group{ctx: ctx}
}

func ReadTask() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	g := WithContext(ctx)
	var doneErr error
	fmt.Println("start", time.Now())
	g.Go(func(ctx context.Context) error {
		go func() {
			time.Sleep(5 * time.Second) // 模拟任务耗时5s
			fmt.Println("模拟任务耗时5s", time.Now())
		}()
		select {
		case <-ctx.Done():
			doneErr = ctx.Err()
		}
		return doneErr
	})

	err := g.Wait()
	if err != nil {
		fmt.Println("err", err)
		fmt.Println("err time", time.Now())
		return
	}
	if doneErr != context.Canceled {
		fmt.Println("Canceled")
	}
	fmt.Println("done time", time.Now())
}

func (g *Group) do(f func(ctx context.Context) error) {
	ctx := g.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	var err error
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)           // 65536
			buf = buf[:runtime.Stack(buf, false)] // 只获取当前堆栈信息
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err // 如果出现err它会cancel掉未执行的goroutine
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
		g.wg.Done() // 如果发生错误了就执行done方法。
	}()
	err = f(ctx)
}

// 这段Go代码是用于处理Go语言中的panic异常的。在Go语言中，当程序发生panic异常时，程序会立即停止执行并打印出一个stack trace信息。
//这段代码使用了Go语言中的recover函数来捕获这个panic异常，然后通过打印一个带有stack trace信息的错误消息来处理这个异常。
//
//具体来说，这段代码使用了Go语言中的短变量声明（short variable declaration）方式将recover函数返回的值赋值给了变量r。
//然后，它创建了一个长度为64KB的byte数组，并将其缩小到空的slice。接下来，它使用Go语言中的runtime包的Stack函数来获取一个stack trace信息，并将其存储在buf中。
//最后，它使用fmt包的Errorf函数将panic信息和stack trace信息组合成一个错误消息，并将其返回。
//
//需要注意的是，这段代码通常是在使用errgroup或其他类似的并发控制库时使用的，用于避免一个goroutine的panic导致整个程序崩溃。

// 按理说 批量请求接口，组装请求成功的数据。
