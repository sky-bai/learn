package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// 执行一个子goroutine去执行一个任务 耗时10s
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				println("任务结束")
				return
			default:
				println("hello")
				time.Sleep(3 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
}
