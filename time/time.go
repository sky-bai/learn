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

	timeSub()

	ti, err := time.Parse(time.RFC3339, "2022-12-07T16:00:05.000Z")
	fmt.Println(ti, err)
	if ti.Minute() == 0 {
		fmt.Println("true")
	}
	fmt.Println("--- second", ti.Second())

	fmt.Println(time.Now().Format("2006-0102-150405"))

}

func timeSub() {
	fmt.Println(time.Now().AddDate(0, 0, -1).Sub(time.Now()))

	// 给当前时间减去两小时
	fmt.Println(time.Now().Add(-2 * time.Hour))
}
