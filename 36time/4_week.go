package main

import (
	"fmt"
	"time"
)

func TimeSequential(date time.Time, typ string) int64 {
	var r int64
	switch typ {
	case "WEEK":
		// 计算从 1970 年 1 月 1 日至今的周数
		r = int64(date.Sub(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)).Hours()/(24*7)) + 1
	}
	return r
}

func main() {
	t := time.Now().UTC().Round(time.Second)
	fmt.Println(TimeSequential(t, "WEEK"))

	t1 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	fmt.Println(TimeSequential(t1, "WEEK"))

	milliTime := time.Now().UTC().Round(time.Second).UnixMilli()
	fmt.Println("milliTime", milliTime)

	// 时间戳转换成北京时间
	// 1. 传入时间戳
	// 2. 传入时间戳转换成北京时间
	// 3. 传入时间戳转换成北京时间然后转成utc时间
	mill := 1670392298000
	// 将mill 转换成北京时间
	day := time.Unix(int64(mill/1000), 0)
	fmt.Println("day", day)

	milliTime1 := day.UTC().Round(time.Second).UnixMilli()
	fmt.Println("milliTime1", milliTime1)

}
