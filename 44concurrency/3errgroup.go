package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {

	g, _ := errgroup.WithContext(context.Background())

	// 调用广告
	g.Go(func() error {
		return errors.New("test")
	})
	// 调用AI
	g.Go(func() error {
		return errors.New("test")
	})
	// 调用运营平台
	g.Go(func() error {
		return errors.New("test")
	})
	// 我们把一个复杂的任务，尤其是依赖多个微服务rpc需要聚合数据的任务，分解为依赖
	//和并行，依赖的意思为:需要上游a的数据才能访问下游b的数据进行组合。但是并行
	//的意思为:分解为多个小任务并行执行，最终等全部执行完毕。
	err := g.Wait()
	fmt.Println(err)
}
