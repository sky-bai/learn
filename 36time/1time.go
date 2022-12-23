package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

func main() {

	str := "2022-11-11 11:11:11"

	loc, _ := time.LoadLocation("Local")

	the_time, _ := time.ParseInLocation("2006-01-02 15:04:05 ", str, loc)
	fmt.Println("------", the_time)
	fmt.Println("11111", now.With(the_time).Year())
	currentTime := time.Now()
	yesterTime := currentTime.Add(-24 * time.Hour)
	fmt.Println(yesterTime)

	utcTime := time.Now().UTC().Round(time.Second)
	fmt.Println("utcTime", utcTime)

}

// Go的time. Now()会设置时间为本地时区，转字符串时，如果有时区就会加上+8的这样的字样，
// 如果是UTC则就不会加。反序列化回来的时候，有时区就会设置，没有时区则设置为nil，
// 这样就导致结构体不一样。解决方案是 time. Now().UTC()，这样在结构中和序列化时时区都为空……

// Go的时间还会在序列化时加上m=+0.002691802，这个叫单调时间，
// 主要是为了防止我们电脑时间跟NTP同步时来回跳的问题，保证时间总是向前移动，
// 并且不会受到导致时间跳跃的变化的影响。但是，这个会导致序列化和反序列化后时间对象不一致，
// 要解决方案是：time. Now().UTC().Round(...) 。
