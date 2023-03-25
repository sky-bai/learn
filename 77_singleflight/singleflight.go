package main

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

// errGoExit indicates the runtime.GoExit was called in  runtime.GoExit 这个函数是用来干什么的
// the user given function.
var errGoExit = errors.New("runtime.GoExit was called")

func main() {
	var norm bool
	norm = false
	if !norm { // norm 进行取反操作,看看是不是true
		fmt.Println("normal", norm)
	} else {
		fmt.Println("not normal", norm)
	}
}

// A panicError is an arbitrary value recovered from a panic
// with the stack trace during the execution of given function.
type panicError struct {
	value interface{}
	stack []byte
}

// Error implements error interface.
func (p *panicError) Error() string {
	return fmt.Sprintf("%v\n\n%s", p.value, p.stack)
}

func newPanicError(v interface{}) error {
	stack := debug.Stack()

	// The first line of the stack trace is of the form "goroutine N [status]:"
	// but by the time the panic reaches Do the goroutine may no longer exist
	// and its status will have changed. Trim out the misleading line.
	if line := bytes.IndexByte(stack[:], '\n'); line >= 0 {
		stack = stack[line+1:]
	}
	return &panicError{value: v, stack: stack}
}

// call is an in-flight or completed singleflight.Do call
type call struct {
	wg sync.WaitGroup

	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
	val interface{}
	err error

	// These fields are read and written with the singleflight
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
	dups  int
	chans []chan<- Result
}

// Group 结构体由一个互斥锁和一个 map 组成，可以看到注释 map 是懒加载的，所以 Group 只要声明就可以使用，不用进行额外的初始化零值就可以直接使用。
// call 保存了当前调用对应的信息，map 的键就是我们调用 Do 方法传入的 key
// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

// Result holds the results of Do, so they can be passed 保存结果
// on a channel.
type Result struct {
	Val    interface{}
	Err    error
	Shared bool
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	g.mu.Lock()

	// 前面提到的懒加载
	if g.m == nil { // 没有就创建 也就是说一个结构体里面有引用类型的话，就可以这么干
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		// 如果存在就会解锁
		c.dups++
		g.mu.Unlock()

		// 然后等待 WaitGroup 执行完毕，只要一执行完，所有的 wait 都会被唤醒
		c.wg.Wait()

		// 这里区分 panic 错误和 runtime 的错误，避免出现死锁，后面可以看到为什么这么做
		if e, ok := c.err.(*panicError); ok {
			// 如果返回的是 panic 错误，为了避免 channel 死锁，我们需要确保这个 panic 无法被恢复
			panic(e)
		} else if c.err == errGoExit {
			runtime.Goexit()
		}
		return c.val, c.err, true
	}
	//  如果我们没有找到这个 key 就 new call
	c := new(call)
	// 然后调用 waitgroup 这里只有第一次调用会 add 1，其他的都会调用 wait 阻塞掉
	// 所以这要这次调用返回，所有阻塞的调用都会被唤醒
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// 然后我们调用 doCall 去执行
	g.doCall(c, key, fn)
	return c.val, c.err, c.dups > 0
}

// Do chan 和 Do 类似，其实就是一个是同步等待，一个是异步返回，主要实现上就是，如果调用 DoChan 会给 call.chans 添加一个 channel 这样等第一次调用执行完毕之后就会循环向这些 channel 写入数据。

// DoChan is like Do but returns a channel that will receive the
// results when they are ready.
//
// The returned channel will not be closed.
func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result {
	ch := make(chan Result, 1)
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		c.dups++
		c.chans = append(c.chans, ch)
		g.mu.Unlock()
		return ch
	}
	c := &call{chans: []chan<- Result{ch}}
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	go g.doCall(c, key, fn)

	return ch
}

// doCall
// 这个方法种有个技巧值得学习，使用了两个 defer 巧妙的将 runtime 的错误和我们传入 function 的 panic 区别开来避免了由于传入的 function panic 导致的死锁。

// doCall handles the single call for a key.
func (g *Group) doCall(c *call, key string, fn func() (interface{}, error)) {
	normalReturn := false
	recovered := false
	// 第一个 defer 检查 runtime 错误
	// use double-defer to distinguish panic from runtime.Goexit,
	// more details see https://golang.org/cl/134395
	defer func() {
		// 如果既没有正常执行完毕，又没有 recover 那就说明需要直接退出了
		// the given function invoked runtime.Goexit
		if !normalReturn && !recovered {
			c.err = errGoExit
		}

		g.mu.Lock()
		defer g.mu.Unlock()
		c.wg.Done()
		if g.m[key] == c {
			delete(g.m, key)
		}

		if e, ok := c.err.(*panicError); ok {
			// In order to prevent the waiting channels from being blocked forever,
			// needs to ensure that this panic cannot be recovered.
			if len(c.chans) > 0 {
				go panic(e)
				select {} // Keep this goroutine around so that it will appear in the crash dump.
			} else {
				panic(e)
			}
		} else if c.err == errGoExit {
			// Already in the process of goexit, no need to call again
			// 已经准备退出了，也就不用做其他操作了
		} else {
			// Normal return 正常返回
			// 正常情况下向 channel 写入数据
			for _, ch := range c.chans {
				ch <- Result{c.val, c.err, c.dups > 0}
			}
		}
	}()
	// 使用一个匿名函数来执行
	func() {
		defer func() {
			if !normalReturn {
				// 如果 panic 了我们就 recover 掉，然后 new 一个 panic 的错误
				// 后面在上层重新 panic
				// Ideally, we would wait to take a stack trace until we've determined
				// whether this is a panic or a runtime.Goexit.
				//
				// Unfortunately, the only way we can distinguish the two is to see
				// whether the recover stopped the goroutine from terminating, and by
				// the time we know that, the part of the stack trace relevant to the
				// panic has been discarded.
				if r := recover(); r != nil {
					c.err = newPanicError(r)
				}
			}
		}()

		c.val, c.err = fn()
		// 如果 fn 没有 panic 就会执行到这一步，如果 panic 了就不会执行到这一步
		// 所以可以通过这个变量来判断是否 panic 了
		normalReturn = true
	}()
	// 如果 normalReturn 为 false 就表示，我们的 fn panic 了
	// 如果执行到了这一步，也说明我们的 fn  recover 住了，不是直接 runtime exit
	if !normalReturn {
		recovered = true
	}
}

// Forget tells the singleFlight to forget about a key.  Future calls
// to Do for this key will call the function rather than waiting for
// an earlier call to complete.
func (g *Group) Forget(key string) {
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()
}
