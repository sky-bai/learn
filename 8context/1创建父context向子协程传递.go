package main

import (
	"context"
	"time"
)

func main() {
	// 给顶层context创建一个带有完成队列的拷贝
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // done 是一个channel
				return
			default:
				println("work")
				//time.Sleep(time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 当cancel被调用的时候，会发送一个信号给上面的协程，让它退出 会让done的channel收到信号 channel被关闭
	// cancel()方法应该放在子goroutine中，因为子goroutine可以通过ctx.Done()来检测父协程是否已经退出
	time.Sleep(time.Second * 5)
}
