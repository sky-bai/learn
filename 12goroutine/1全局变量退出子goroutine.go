package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exist bool

// break 结束 for、switch 和 select 的代码块
// continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。

func worker() {
	// 子goroutine完成时需要表明该任务已完成
	defer wg.Done()
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exist {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go worker()
	// 父goroutine通知子goroutine退出
	exist = true
	wg.Wait()

	fmt.Println("over")
}
