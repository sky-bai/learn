package main

func main() {

}

// 控制子协程生命周期
// 1.在子goroutine执行过程中，可以通过触发Hook来达到控制子goroutine的目的（通常是取消，即让其停下来）

//WithTimeout 和 WithDeadline 基本上一样，这个表示是超时自动取消，是多少时间后自动取消 Context 的意思。

// 带cancel返回值的Context，一旦cancel被调用，即取消该创建的context
//func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// WithCancel 函数，传递一个父 Context 作为参数，返回子 Context，以及一个取消函数用来取消 Context。

// 带有效期cancel返回值的Context，即必须到达指定时间点调用的cancel方法才会被执行
//func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

// 带超时时间cancel返回值的Context，类似Deadline，前者是时间点，后者为时间间隔
// 相当于WithDeadline(parent, time.Now().Add(timeout)).
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// WithDeadline 函数，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消 Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。

// WithValue 函数和取消 Context 无关，它是为了生成一个绑定了一个键值对数据的 Context，这个绑定的数据可以通过 Context.Value 方法访问到
// func WithValue(parent Context, key, val interface{}) Context
// WithValue 函数和取消 Context 无关，它是为了生成一个绑定了一个键值对数据的 Context，这个绑定的数据可以通过 Context.Value 方法访问到。
// 相当于有这四个方法，就相当于就是要控制子协程，然后可以提前结束子协程的声明周期。

//
// select {
// case <-ctx.Done():
// // 当子goroutine收到取消信号，即做一些相关处理，例如退出Goroutine。
// }
//
