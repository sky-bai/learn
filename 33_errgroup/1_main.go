package main

import (
	"golang.org/x/sync/errgroup"
	"sync"
)

// 了解各个组件如何集成
// WaitGroup 用来解决当前 goroutine 等待多个 goroutine 结束的问题，但是当我们需要获得协程返回的错误，则无能为力。
// 包含了个 WaitGroup 用于同步等待所有 Gorourine 执行

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Wait()

	errGroup := errgroup.Group{}
	errGroup.Wait()

}

type BaseGroup struct {
	cancel func() // 这里保存的是 contex.WithCancel 返回的 cancel 参数，用以取消所有子协程的运行
	// 包含了个 WaitGroup 用于同步等待所有 Gorourine 执行
	wg      sync.WaitGroup
	errOnce sync.Once
	// golang 特有的单例模式，利用原子操作进行锁定值判断只会传递第一个出错的协程的 error
	err error
}

// 并发模式的区分
// 不同场景下的并发模式选择

// 1.对一组操作进行超时限制
// 2.并发处理一批任务

// todo 问题
// 1.cancel 的内部实现 如何取消子协程
// 2.once 的内部实现
