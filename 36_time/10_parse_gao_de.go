package main

import (
	"fmt"
	"time"
)

func main() {

	timeStr := "2024-01-06 13:10:30"
	layout := "2006-01-02 15:04:05"

	// 解析时间字符串
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("解析时间字符串时出现错误:", err)
		return
	}
	// 将时间对象转换成本地时间
	localTime := t.In(time.UTC)
	// 将时间对象转换成时间戳
	timestamp := localTime.Unix()
	fmt.Println("时间戳:", timestamp)
}
