package main

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"runtime"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}, ctx context.Context) {
	time.Sleep(3 * time.Second)

	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
	select {
	case <-ctx.Done():

		fmt.Printf("%d,任务取消", i)
		runtime.Goexit()
	}

	//time.Sleep(10 * time.Second)

}

//func demoFunc() {
//	time.Sleep(10 * time.Millisecond)
//	fmt.Println("Hello World!")
//}

func main() {
	defer ants.Release()

	runTimes := 10
	ctx1 := context.Background()
	var i int32

	// Use the common pool.
	//var wg sync.WaitGroup
	syncCalculateSum := func() {
		myFunc(i, ctx1)
		//wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		//wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	//time.Sleep(4 * time.Second)
	ctx1.Done()
	//wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	//p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
	//	myFunc(i)
	//	wg.Done() // 每个任务完成后，执行done方法。
	//})
	//defer p.Release()
	//// Submit tasks one by one.
	//for i := 0; i < runTimes; i++ {
	//	wg.Add(1)
	//	_ = p.Invoke(int32(i))
	//}
	//wg.Wait()
	//fmt.Printf("running goroutines: %d\n", p.Running())
	//fmt.Printf("finish all tasks, result is %d\n", sum)
}
