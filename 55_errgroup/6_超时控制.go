package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	fmt.Println("time", time.Now())
	// 设置上下文的超时为1秒钟
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// 模拟一个耗时2秒的操作
		time.Sleep(2 * time.Second)
		return errors.New("task1 failed")
	})
	g.Go(func() error {
		// 模拟一个耗时3秒的操作
		time.Sleep(3 * time.Second)
		return errors.New("task2 failed")
	})

	var err error

	// 等待所有任务完成
	if err = g.Wait(); err != nil {
		fmt.Println(time.Now())
		fmt.Printf("Error: %v\n", err)

	}

	// 如果上下文被取消，则在这里处理取消操作
	select {
	case <-ctx.Done():
		fmt.Printf("Canceled: %v\n", ctx.Err())
	default:
		fmt.Println("Success")
	}
	fmt.Println("time", time.Now())
}
