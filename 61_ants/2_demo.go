package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"os"
	"sync"
	"time"
)

func demoFunc1() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello World!")
}

func wrapper(i int, wg *sync.WaitGroup) func() {
	return func() {
		// 这里一定要保证任务要完成,否则wg.Wait()会永远阻塞
		defer wg.Done()
		if i%2 == 0 {
			panic(fmt.Sprintf("panic from task:%d", i))
		}
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Printf("hello from task:%d\n", i)

	}
}

func panicHandler(err interface{}) {
	fmt.Fprintln(os.Stderr, err)
}

func main() {
	p, _ := ants.NewPool(2, ants.WithPanicHandler(panicHandler))
	defer p.Release()

	runTimes := 10 // 执行次数

	// 1.不知道任务类型 直接使用 ants.Submit
	// Use the common pool. 使用一组并发请求
	var wg sync.WaitGroup
	//syncCalculateSum := func() { // 需要执行的函数 如果用wg的话 需要在每个执行的任务中添加wg.Done()来表示执行完成
	//	defer wg.Done()
	//	demoFunc1()
	//
	//}
	// 模拟1000次请求
	for i := 0; i < runTimes; i++ {
		wg.Add(1) // 计数器添加次数
		j := 0
		if i == 0 {
			j = 10
		} else {
			j = 1
		}
		err := p.Submit(wrapper(j, &wg))
		if err != nil {
			fmt.Println("err:", err)
		}
	}

	wg.Wait() // 使用wg的话 可以等待所有任务执行完成

}
