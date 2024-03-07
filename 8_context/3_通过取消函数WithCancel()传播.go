package main

import (
	"context"
	"fmt"
	"time"
)

func someHandler() {
	// 通过调用Background创建父节点parentContext
	parentContext := context.Background()
	// 创建通过WithCancel方法继承Background的子节点childContext
	childContext, cancel := context.WithCancel(parentContext)
	go doSth(childContext, "child [1]")
	go doSth(childContext, "child [2]")

	//模拟程序运行 - Sleep 3秒
	time.Sleep(3 * time.Second) // 3s之后进行关闭
	cancel()
	//等待取消信号能被子Context读到
	time.Sleep(2 * time.Second)
}

// 每1秒work一下，同时会判断ctx是否被取消，如果是就退出
func doSth(ctx context.Context, name string) {
	var i = 1
	for {
		//模拟具体work - Sleep 1秒
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Printf("%s done!\n", name)
			return
		default:
			fmt.Printf("%s had worked %d seconds \n", name, i)
		}
		i++
	}
}

func main() {
	fmt.Println("start.")
	someHandler()
	fmt.Println("end.")
}
