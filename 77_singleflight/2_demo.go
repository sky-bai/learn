package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type Result1 string

func find(ctx context.Context, query string) (Result1, error) {
	return Result1(fmt.Sprintf("result for %q", query)), nil
}

func main() {
	DoDemo()
	DoChanDemo()
}

func DoDemo() {
	var g Group
	const n = 5
	waited := int32(n)
	done := make(chan struct{})
	key := "https://weibo.com/1227368500/H3GIgngon"
	for i := 0; i < n; i++ {
		go func(j int) {
			v, _, shared := g.Do(key, func() (interface{}, error) {
				ret, err := find(context.Background(), key)
				return ret, err
			})
			if atomic.AddInt32(&waited, -1) == 0 {
				close(done)
			}
			fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
		}(i)
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		fmt.Println("Do hangs")
	}
}

func DoChanDemo() {

}

// 阻塞读是什么？
// 由于 singleFlight 是以阻塞读的方式来控制向下游请求的并发量，在第一个下游请求没有返回之前，所有请求都将被阻塞。 多次请求,一个人执行，但是这次执行耗时较长，那么后续请求都会被阻塞，直到第一个请求返回。
// 假设服务正常情况下处理能力为 1W QPS，每次请求会发起 3 次 下游调用，其中一个下游调用使用 singleflight.md 获取控制并发获取数据，请求超时时间为3S。那么在出现请求超时的情况下，会出现以下几个问题：
//
// 协程暴增，最小协程数为3W（1 W/S * 3S）
// 内存暴涨，内存总大小为：协程内存大小 + 1W/S * 3S *（3+1）次 * （请求包+响应包）大小
// 大量超时报错：1W/S * 3S
// 后续请求耗时增加（调度等待）
// 如果类似问题出现在重要程度高的接口上，例如：读取游戏配置、获取博主信息 等关键接口，那么问题将是非常致命的。出现该情况的根本原因有以下两点：
//
// 1.阻塞读：缺少超时控制，难以快速失败
// 2.单并发：控制了并发量，但牺牲了成功率，减少并发，牺牲成功率
// 那么如何应对以上问题呢？
