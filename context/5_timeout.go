package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个上下文对象，并设置超时时间为 1 秒
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 启动一个 goroutine 执行一些工作，可能会在超时时间到达前完成，也可能超时
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("work done")
	}()

	// 在上下文对象被取消前，阻塞等待工作完成或者超时
	select {
	case <-ctx.Done(): // 1s 中会执行done操作
		fmt.Println(ctx.Err())
	}
}

// 在上面的代码中，我们使用了 defer cancel() 语句来确保在 main 函数退出前调用 cancel() 函数取消上下文对象。
// 这是因为如果上下文对象在超时时间到达前被取消，它会通知其它正在等待它的 goroutine 退出，以避免资源泄漏和其它问题。
// 因此，我们需要确保在 main 函数退出前，取消上下文对象，以确保 goroutine 能够正常退出。
//
// defer 语句用于延迟函数的执行，在函数返回前执行。在这个例子中，我们使用 defer 语句延迟执行 cancel() 函数，确保在 main 函数退出前，取消上下文对象。
//
// 需要注意的是，defer 语句是按照定义的顺序逆序执行的。因此，在多个 defer 语句存在的情况下，最后定义的 defer 语句会最先执行，而最先定义的 defer 语句会最后执行。
// 在这个例子中，我们定义了一个 defer cancel() 语句，它会在最后执行，确保在函数返回前取消上下文对象。
