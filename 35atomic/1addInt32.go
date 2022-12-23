package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Add() {
	var x int32
	// 多协程执行并发累计变量
	var wg sync.WaitGroup

	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("完成", x)
}
