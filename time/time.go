package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"sync/atomic"
	"time"
)

func main() {
	oneMonth := time.Now().AddDate(0, -1, 0)
	startDay := now.New(oneMonth).BeginningOfMonth()
	endDay := now.New(oneMonth).EndOfMonth()
	fmt.Println("", startDay)
	fmt.Println("", endDay)

	var TickerCount int64
	fmt.Println("", TickerCount)

	atomic.LoadInt64(&TickerCount) //获取数据 在这里读取value的值的同时，当前计算机中的任何CPU都不会进行其它的针对此值的读或写操作。
	fmt.Println("", TickerCount)

	atomic.StoreInt64(&TickerCount, 1)
	fmt.Println("", TickerCount)
}
