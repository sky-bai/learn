package main

import (
	"fmt"
	"time"
)

func main() {
	timeParam := "2022-12-06T16:00:00.000Z"
	// 将time转成time.Time类型
	loc, _ := time.LoadLocation("Asia/Shanghai")

	the_time, _ := time.ParseInLocation("2006-01-02 15:04:05 ", timeParam, loc)
	fmt.Println("------", the_time)

	fmt.Println("day", time.Now().Day())

}
