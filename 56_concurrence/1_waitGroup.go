package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println(time.Now())
	// 声明一个等待组
	var wg sync.WaitGroup
	// 准备一系列的网站地址
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}
	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时, 将等待组增加1
		wg.Add(1)
		// 开启一个并发
		go func(url string) {
			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()
			// 使用http访问提供的地址
			_, err := http.Get(url)
			// 访问完成后, 打印地址和可能发生的错误
			fmt.Println(url, err)
			// 通过参数传递url地址
		}(url)
	}
	// 等待所有的任务完成
	wg.Wait()
	fmt.Println("over")
	fmt.Println(time.Now())
}

// 下载多个文件
// 当第一个下载好了之后就开始进行糊化 糊化的时候就只要两个地址 一个是获取地址 一个是保存地址
// 然后所有糊化完成就结束

// 使用方法
// 任务开始时调用add方法
// 任务完成时调用done方法
// 等待所有任务完成时调用wait方法

//Go语言中的sync包中提供了一个WaitGroup类型，用于协调多个goroutine的执行。WaitGroup可以用于等待一组goroutine执行完毕，然后再继续执行接下来的代码，它可以确保所有的goroutine都已经执行完毕，避免了可能的竞态条件和数据竞争。
//下面是一些使用WaitGroup的场景：
//在并发地执行多个goroutine时，需要确保所有goroutine都已经执行完毕再执行接下来的代码。
//在一个goroutine中，需要等待其他多个goroutine执行完毕，再继续执行接下来的代码。
//在一个goroutine中，需要等待一组子goroutine中的任意一个goroutine执行完毕，再继续执行接下来的代码。
//需要注意的是，WaitGroup只是一个计数器，它不能保证goroutine的执行顺序，也不能确保所有的goroutine都已经执行成功。如果需要控制goroutine的执行顺序或处理goroutine的错误，可以使用其他的并发控制机制，例如Channel或者Kratos框架中提供的errgroup。
// 它的缺点是不能组装有效数据
// 无法进行超时控制

// 需求1 每一个任务有超时控制，如果超时，就不再等待了
// 需求2 每一个任务有重试机制，如果失败，就重试
// 需求3 每一个任务有并发控制，最多只能有3个任务在执行
// 需求4 控制整体的超时时间，如果时间到了，组装请求成功的任务，任务失败或者超时就返回默认的值

// 1.之前的waitGroup是没有超时控制的，如果任务执行时间过长，就会一直等待下去
