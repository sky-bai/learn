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

var key = "cfStat_track_flow_" + "channel" + "20231206"
var k1 = "cfStat_track_num_" + "channel" + "20231206"

// Incr 是一个简单的累加器
type Incr struct {
	keys []string

	keyAndCount map[string]uint64
}

// NewIncr 创建一个新的 Incr 实例
func NewIncr(keys []string) *Incr {
	return &Incr{
		keys: keys,
	}
}

// Incr 对给定的 key 进行累加操作
func (si *Incr) Incr(key, val uint64) {
	atomic.AddUint64(&key, val)
}

// TakeoutToRedis 获取给定 key 的当前累加值并重置为0
func (si *Incr) TakeoutToRedis(key uint64) uint64 {
	value := atomic.SwapUint64(&key, 0)
	// redis.incr(key, value)
	return value
}

// 要做到每日更新
