package main

import (
	"fmt"
	"time"
)

func main() {

	//str := "2024-01-06 13:10:30"
	str1 := ""
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05 ", str1, loc)
	fmt.Println("parseInLocation", theTime)

	fmt.Println("1111", theTime.UnixNano()/1e6)
	fmt.Println("1111", theTime.UnixMilli())

	//fmt.Println("11111", now.With(theTime).Year())
	//currentTime := time.Now()
	//yesterTime := currentTime.Add(-24 * time.Hour)
	//fmt.Println(yesterTime)

	utcTime := time.Now().UTC().Round(time.Second)
	fmt.Println("utcTime", utcTime)

	fmt.Println("time.now.string", time.Now().String())
	fmt.Println("time.now.utc", time.Now().UTC().String())

	fmt.Println(time.Now().UnixNano() / 1e6)

	initialTimestamp := time.Unix(0, 0).UnixMilli()
	fmt.Println("最初的时间戳:", initialTimestamp)

	currentTime := time.Now()
	fmt.Println("ooo", currentTime)

	// 计算前一分钟的时间戳
	twoMinuteAgo := currentTime.Add(-time.Minute * 2)
	fmt.Println("ooo", twoMinuteAgo)

}

// Go的time.Now()会设置时间为本地时区，转字符串时，如果有时区就会加上+8的这样的字样,如果是UTC则就不会加。
// 反序列化回来的时候，有时区就会设置，没有时区则设置为nil，
// 这样就导致结构体不一样。解决方案是 time. Now().UTC()，这样在结构中和序列化时时区都为空……

// Go的时间还会在序列化时加上m=+0.002691802，这个叫单调时间，
// 主要是为了防止我们电脑时间跟NTP同步时来回跳的问题，保证时间总是向前移动，
// 并且不会受到导致时间跳跃的变化的影响。但是，这个会导致序列化和反序列化后时间对象不一致，
// 要解决方案是：time. Now().UTC().Round(...) 。

// time.now就是加了本地时区，前端显示的时候，要转成utc时间，然后再转成本地时区的时间
// 如果我直接就设置为utc时间，那么前端显示的时候，就不需要转了
