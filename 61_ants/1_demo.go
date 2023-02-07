package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

// 累加传入的值
func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

// 每十毫秒执行一次打印任务
func demoFunc() {
	time.Sleep(10 * time.Millisecond) // 毫秒
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	runTimes := 1000 // 执行次数

	// 1.不知道任务类型 直接使用 ants.Submit
	// Use the common pool. 使用一组并发请求
	var wg sync.WaitGroup
	syncCalculateSum := func() { // 需要执行的函数 如果用wg的话 需要在每个执行的任务中添加wg.Done()来表示执行完成
		demoFunc()
		wg.Done()
	}
	// 模拟1000次请求
	for i := 0; i < runTimes; i++ {
		wg.Add(1) // 计数器添加次数
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait() // 使用wg的话 可以等待所有任务执行完成
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// 2.已知任务类型,只需要提供invoke函数供上层调用
	// Use the pool with a function,  全部添加到协程池
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

// 添加定时任务的时候不能通过协程
// 时间轮删除任务在设置任务之前 然后还是会执行一次
// 调用三方组件 查看耗时
// 先从redis取 取不到再从mysql取
