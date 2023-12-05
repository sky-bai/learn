package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var x1 uint64

func main() {
	atomic.AddUint64(&x1, 1)
	fmt.Println(x1)
	atomic.AddUint64(&x1, 1)

	get := atomic.SwapUint64(&x1, 0)
	fmt.Println(get)
	fmt.Println(x1)

	// 等待中断信号
	quit := make(chan os.Signal)

	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("程序退出")

	fmt.Println(x1)

}
