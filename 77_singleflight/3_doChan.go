package main

import (
	"context"
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {

}

// 普通调用方法
func callFunc(i int) (int, error) {
	time.Sleep(500 * time.Millisecond)
	return i, nil
}

// 使用singleflight

// 1. 定义全局变量
var sf singleflight.Group

func callFuncBySF(key string, i int) (int, error) {
	// 2. 调用sf.Do方法
	value, err, _ := sf.Do(key, func() (interface{}, error) {
		return callFunc(i)
	})
	res, _ := value.(int)
	return res, err
}

// CtrTimeout 使用DoChan进行超时控制
func CtrTimeout(ctx context.Context, req interface{}) {
	ch := sf.DoChan(key, func() (interface{}, error) {
		return call(ctx, req)
	})

	select {
	case <-time.After(500 * time.Millisecond):
		return
	case <-ctx.Done():
		return
	case ret := <-ch:
		go handle(ret)
	}
}

// 另外启用协程定时删除key，提高请求下游次数，提高成功率
func CtrRate(ctx context.Context, req interface{}) {
	res, _, shared := g.Do(key, func() (interface{}, error) {
		// 另外其一个goroutine，等待一段时间后，删除key
		// 删除key后的调用，会重新执行Do
		go func() {
			time.Sleep(10 * time.Millisecond)
			g.Forget(key)
		}()

		return call(ctx, req)
	})

	handle(res)
}
