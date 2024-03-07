package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	foo(ctx)
}

func foo(ctx context.Context) {
	stop := make(chan struct{}, 0)
	go func() {
		slow()
		// 在slow函数后给一个通道赋值 表示该函数执行完成
		//stop <- struct{}{}
		close(stop)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("context timeout")
		return
	case <-stop:
		fmt.Println("slow done")
		return
	}

}

func slow() {
	n := rand.Intn(3)
	fmt.Printf("sleep  %d s\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}

// 只能编辑foo函数
// foo 必须调用slow函数
// foo函数必须在ctx超时后立刻返回
