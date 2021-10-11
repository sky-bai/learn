package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// 开十个读 每一毫秒读一次
	var count Counter
	for i := 0; i < 10; i++ {
		go func() {
			count.GetCount()
			fmt.Println("读到的count", count)
			time.Sleep(time.Millisecond)
		}()
	}

	// 一个写 每一秒写一次
	for {
		count.Incr()
		fmt.Println("增加后的count", count)
		time.Sleep(time.Second)
	}

}

type Counter struct {
	mu    sync.RWMutex
	count uint64
}

// count 加一
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 读count
func (c *Counter) GetCount() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
