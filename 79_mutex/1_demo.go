package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool, 1) // 构造一个是否成功的chan 信号
	var mu sync.Mutex
	i, j := 0, 0

	// goroutine 1
	go func() { // 开一个协程监听事件 go func select
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock() // 唯一执行一段代码逻辑
				i++
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	// goroutine 2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		j++
		mu.Unlock()
	}
	done <- true
	fmt.Println("i", i)
	fmt.Println("j", j)
}
