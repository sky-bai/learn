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
	// 使用 atomic 包提供的原子操作来访问和修改共享变量可以确保在多个 goroutine 中的安全性。

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("完成", x)
}

// 推荐文章: https://www.codenong.com/js2c7797df9b2b/
// 兼容不同的计算机体系结构，从语言层面提供一个统一的函数。
//例如在现在32位/64位 RISC计算机体系结构下面，一条load/store指令就能读取32位/64位的数据，那么CPU指令级别确实能保证int32的读取完整性，也就LoadInt32/StoreInt32失去存在的必要性。
//但是不同的计算机体系结构，不一定能提供这种指令，举个例子部分老式CPU可能需要4条指令，一次读取一个字节，这种情况下就必须保证读写的原子性。
//还有啊在32位计算机体系结构下，读写一个64位的数据，必然需要两条指令，这时LoadInt64/StoreInt64就需要了。
