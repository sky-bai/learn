package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var TickerCount int64
	TickerCount = 1
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(20 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			tickerCount := atomic.LoadInt64(&TickerCount) //获取数据
			fmt.Println("TickerCount:", tickerCount)
			fmt.Println("Current time: ", t)
		}
	}

	// 这个select 监控两个channel
	// ticker 会创建一个channel 然后会根据时间间隔去往channel里面写入数据
}
