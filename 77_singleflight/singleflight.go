package main

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

// errGoExit indicates the 2_runtime.GoExit was called in  2_runtime.GoExit 这个函数是用来干什么的
// the user given function.
var errGoExit = errors.New("2_runtime.GoExit was called")

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

// newPanicError creates a panicError with the given value and stack trace.
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

// call is an in-flight or completed singleflight.md.Do call
type call struct {
	wg sync.WaitGroup // 存储返回值，在wg done之前只会写入一次 只会在第一次请求完成时才会done。

	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
	val interface{}
	err error // 存储返回的错误信息 存储实际操作报错的error

	// These fields are read and written with the singleFlight.md
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
	dups  int             // 统计相同请求的次数，在wg done之前写入
	chans []chan<- Result // 使用DoChan方法使用，用channel进行通知 多个可读的channel
}

// Group 结构体由一个互斥锁和一个 map 组成，可以看到注释 map 是懒加载的，所以 Group 只要声明就可以使用，不用进行额外的初始化零值就可以直接使用。
// call 保存了当前调用对应的信息，map 的键就是我们调用 Do 方法传入的 key
// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m // 互斥锁，保证并发安全
	m  map[string]*call // lazily initialized  保存key对应的函数执行过程和结果的变量。  // 存储相同的请求，key是相同的请求，value保存调用信息。
	// 关于请求和调用函数的map映射表
}

// Result holds the results of Do, so they can be passed 保存结果
// on a channel.
type Result struct { // DoChan方法时使用
	Val    interface{} // 存储返回值
	Err    error       // 存储返回的错误信息
	Shared bool        // 标示结果是否是共享结果
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
// 入参：key：标识相同请求，fn：要执行的函数
// 返回值：v: 返回结果 err: 执行的函数错误信息 shard: 是否是共享结果
func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	g.mu.Lock()

	// 前面提到的懒加载
	if g.m == nil { // 没有就创建 也就是说一个结构体里面有引用类型的话，就可以这么干
		g.m = make(map[string]*call)
	}
	// 判断是否有相同请求
	if c, ok := g.m[key]; ok {
		// 如果存在就会解锁  // 相同请求次数+1
		c.dups++
		// 解锁就好了，只需要等待执行结果了，不会有写入操作了
		g.mu.Unlock()

		// 然后等待 WaitGroup 执行完毕，只要一执行完，所有的 wait 都会被唤醒
		c.wg.Wait()

		// 这里区分 panic 错误和 2_runtime 的错误，避免出现死锁，后面可以看到为什么这么做
		if e, ok := c.err.(*panicError); ok { // 如果是
			// 如果返回的是 panic 错误，为了避免 channel 死锁，我们需要确保这个 panic 无法被恢复
			panic(e)
		} else if c.err == errGoExit { // 什么时候会出现需要退出的error
			/**
			GoExit
			调用runtime.goExit()将立即终止当前goroutine执行，调度器
			确保所有已注册defer延迟调度被执行。
			*/
			runtime.Goexit() // 2_runtime.GoExit函数在终止调用它的Goroutine的运行之前会先执行该Goroutine中还没有执行的defer语句。
		}
		return c.val, c.err, true
	}
	//  如果我们没有找到这个 key 就 new call  // 之前没有这个请求，则需要new一个指针类型
	c := new(call)
	// 然后调用 waitgroup 这里只有第一次调用会 add 1，其他的都会调用 wait 阻塞掉
	// 所以这要这次调用返回，所有阻塞的调用都会被唤醒
	c.wg.Add(1)   // sync.waitgroup的用法，只有一个请求运行，其他请求等待，所以只需要add(1)
	g.m[key] = c  // m赋值 赋值结果 删除的时候会判断c是否相同
	g.mu.Unlock() // 没有写入操作了，解锁即可

	// 然后我们调用 doCall 去执行
	g.doCall(c, key, fn) // 唯一的请求该去执行函数了 函数作为函数参数 函数也就被当成变量进行传入  那么这个特殊的变量在实例它的时候就会被执行 那么函数被创建为一个变量的时候就可以被使用了 那么这个和直接调用有什么区别昵
	return c.val, c.err, c.dups > 0
}

// 对新请求的处理 对于新请求，我们会调用 doCall 方法，这个方法的主要作用就是执行 fn 函数，然后把结果写入 call 结构体中，然后唤醒所有等待的调用。
// 对重复请求的处理

// Do 同步等待 Do chan 异步返回
// Do chan 和 Do 类似，其实就是一个是同步等待，一个是异步返回，主要实现上就是，如果调用 DoChan 会给 call.chans 添加一个 channel 这样等第一次调用执行完毕之后就会循环向这些 channel 写入数据。

// 阻塞读
// 作为 Do() 的替代函数，singleFlight 提供了 DoChan（）。两者实现上完全一样，不同的是，DoChan() 通过 channel 返回结果。因此可以使用 select 语句实现超时控制

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
		c.dups++                      // 加锁和map判断是否存在 然后把dups++
		c.chans = append(c.chans, ch) // 这里是唯一的区别，如果是重复请求，就把 channel 添加到 call.chans 中 重复请求是一组切片管道
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
// 这里使用了两个 defer 巧妙的将 2_runtime 的错误和我们传入 function 的 panic 区别开来避免了由于传入的 function panic 导致的死锁。
// 第一个defer在何时使用，第二个defer在何时使用

// doCall handles the single call for a key.
func (g *Group) doCall(c *call, key string, fn func() (interface{}, error)) { // 对实际执行操作的结果进行处理

	// 1.先用匿名函数执行实际操作函数 做好recover操作
	// 2.函数执行完就删除key
	normalReturn := false // 标识是否正常返回
	recovered := false    // 标识别是否发生panic
	// 第一个 defer 检查 2_runtime 错误
	// use double-defer to distinguish panic from 2_runtime.Goexit,
	// more details see https://golang.org/cl/134395
	defer func() { // 抓住本函数错误
		// 如果既没有正常执行完毕，又没有 recover 那就说明需要直接退出了
		// the given function invoked 2_runtime.Goexit // 通过这个来判断是否是runtime导致直接退出了
		if !normalReturn && !recovered { // 没有正常返回和没有recover住
			c.err = errGoExit // 返回runtime错误信息
		}

		g.mu.Lock()
		defer g.mu.Unlock()
		c.wg.Done()
		if g.m[key] == c { // 防止重复删除key    结果相同才删除key
			delete(g.m, key) // 当有一个请求执行完毕，就从map中删除这个key
		}
		// 检测是否出现了panic错误
		if e, ok := c.err.(*panicError); ok { // 如果是fn实际操作函数的实际错误 就要防止channel出现死锁
			// In order to prevent the waiting channels from being blocked forever,
			// needs to ensure that this panic cannot be recovered.
			if len(c.chans) > 0 { // 为了防止等待通道永远被阻塞，需要确保这种恐慌无法恢复。
				go panic(e) // 开一个协程panic 没有懂这一步
				select {}   // Keep this goroutine around so that it will appear in the crash dump. // 保持住这个goroutine，这样可以将panic写入crash dump
			} else {
				panic(e)
			}
		} else if c.err == errGoExit { // 如果是程序本身runtime时的error 就直接退出
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
	// 使用一个匿名函数来执行 // 使用匿名函数目的是recover住panic，返回信息给上层
	func() {
		defer func() { // 第一个defer处理实际函数执行时可能出现的error
			if !normalReturn {
				// 如果 panic 了我们就 recover 掉，然后 new 一个 panic 的错误
				// 后面在上层重新 panic
				// Ideally, we would wait to take a stack trace until we've determined
				// whether this is a panic or a 2_runtime.Goexit.
				//
				// Unfortunately, the only way we can distinguish the two is to see
				// whether the recover stopped the goroutine from terminating, and by
				// the time we know that, the part of the stack trace relevant to the
				// panic has been discarded.
				if r := recover(); r != nil {
					c.err = newPanicError(r) // 如果不是正常返回，就把panic的错误信息返回给上层 就把fn的错误往上抛
				}
			}
		}()

		c.val, c.err = fn() // 匿名函数调用传入进来的函数 再用defer进行recover 操作
		// 如果 fn 没有 panic 就会执行到这一步，如果 panic 了就不会执行到这一步
		normalReturn = true // 所以可以通过这个变量来判断是否 panic了 如果fn()函数 panic了，那么就不会执行这个赋值操作，就会recover住这个传入函数的error 正常返回修改标识位
	}()
	// 如果 normalReturn 为 false 就表示我们的 fn panic 了
	// 如果执行到了这一步，也说明我们的 fn  recover 住了，不是直接 2_runtime exit
	if !normalReturn {
		recovered = true // 这里说明没有正常返回，但是recover住了
	}
}

// do方法整体逻辑
// 解决什么样的问题

// Forget tells the singleFlight to forget about a key.  Future calls
// to Do for this key will call the function rather than waiting for
// an earlier call to complete.
func (g *Group) Forget(key string) {
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()
}

// Forget方法的作用是让singleFlight加锁操作忘记某个key，这样后续对这个key的请求会重新执行fn函数，而不是等待之前的请求完成。

// 能够在抑制对下游的多次重复请求
// waitGroup 可以针对不同类型进行管理，这里的singleFlight是通过key进行所有请求的分组管理
// 处理的窗口期是 也就是说当前周期进行归并的请求和第一个请求开始到结束时间间隔区间来的请求
// 在那里删除key？ 判断结果的指针是否相同，如果相同才能删除
// 用go语言的并发编程去解决之前的问题
// 什么时候会出现channel 死锁？

// 并发编程singleFlight的实现

// 在简历上描写对并发编程的实现可以突出你在Go语言中使用了以下关键技术：WaitGroup、ErrGroup、Singleflight和Sync.Pool。以下是一种可能的描述方式：
//
//"在我的学习过程中，我积极探索并发编程的实现，并熟练应用于Go语言。我深入了解了Go语言中的几个关键概念和技术，包括WaitGroup、ErrGroup、Singleflight和Sync.Pool。
//
//通过使用WaitGroup，我能够有效地等待一组并发任务的完成，以确保在继续执行之前所有任务都已完成。这使我能够编写高效的并发代码，提高了程序的性能和响应能力。
//
//ErrGroup是一个强大的并发控制工具，它允许我同时启动多个并发任务，并能够捕获和处理其中任何一个任务发生的错误。这使我能够更好地管理和监控并发任务的执行状态，并做出相应的处理。
//
//SingleFlight是一个有用的并发模式，它在多个并发请求对同一个资源进行访问时，只执行一次实际的请求，并将结果返回给所有的请求方。我成功地利用Singleflight来避免重复的计算和资源浪费，提高了程序的效率。
//
//另外，我还深入了解了Sync.Pool，这是一个对象池的实现，用于复用临时对象，减少内存分配和垃圾回收的压力。通过合理地使用Sync.Pool，我能够降低程序的内存开销，并提高程序的性能和稳定性。
//
//通过运用这些并发编程技术，我能够编写出高效、可伸缩且线程安全的Go语言程序，提高了系统的并发处理能力和资源利用率。"

// 能再简短一点吗，突出我的能力，不要太多的技术细节
// 还有atomic原子包的使用 他对应的是那种能力？任何技术的使用场景
// 需要让别人知道你干了什么 go语言的性能分析该如何做？

// 一定要给别人将清楚
// 把需要阐述的东西都写出来
// 玩转github actions
// 玩转channel 各种使用场景 第三方库的使用场景
// 性能优化离不开的一些套路：异步、去锁、复用、零拷贝、批量等
// go语言去锁的一些套路
// go语言复用的一些套路
// 对于tcp服务进行指标监控，能统计到那些有用的数据昵？

// 现在我这家公司是物联网公司，我是服务端go语言，设备端是c++，之前双方通信是服务端使用node.js通过tcp协议与设备进行通信，我现在使用go语言重写了node.js的通信模块，并加入了时间轮算法去解决链接的超时问题，第二个是重写整个node.js代码，使其变得可扩展性，
// 使用prometheus 对服务进行指标检测，指标有连接数，请求耗时。我应该在简历上如何描述昵，简短一点

// 在简历中简洁地描述你在物联网公司中的工作和贡献，可以如下所示：
//
//"在物联网公司，作为服务端Go语言工程师，负责与C++设备端进行通信的重要模块。我重写了原先使用Node.js的通信模块，将其改写为高效的Go语言实现，并引入时间轮算法解决了链接超时问题。
// 此外，我还全面重写了Node.js代码，使其具备出色的可扩展性和灵活性。通过这些改进，提升了通信效率和系统的可靠性，为公司的物联网产品提供了更优质的服务。"
//

// 现在我有一个tcp的go语言服务，每个连接的Id一般是用什么来表示？为什么要用这个来表示？有什么好处？
// 熟悉算法与数据结构

// tcp服务解决的问题有什么
