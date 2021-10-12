package main

import (
	"fmt"
	"sync"
	"time"
)

// web request 每一个请求都有一个对应的goroutine去处理
// 但是一个请求里面又包含多个操作 比如 访问数据库 和 RPC 服务 这时候又会开启多个goroutine去处理任务
// 这多个goroutine的一个goroutine请求超时的时候需要将其他goroutine进行退出

var wg sync.WaitGroup

func main() {
	wg.Add(1) // 表示我要开启一个goroutine
	go worker()
	wg.Wait()
}

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	wg.Done()
	fmt.Println("over")
}

// 如何结束worker里面的子goroutine
