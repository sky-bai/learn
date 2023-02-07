package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println(4 % 2) // 求出余数
	fmt.Println(4 / 2) // 求出商

	fmt.Println(1675305202965 / 1000) // 求出商

	usedTime, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(time.Now().UnixMilli())/1000/60), 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("---", usedTime)
	// 获取两个小时一分钟前的时间戳
	time.Now().Add(-time.Hour * 2).Add(-time.Minute * 1).Add(-time.Millisecond * 1).UnixMilli()

	fmt.Println("+++", fmt.Sprintf("%.2f", float64(time.Now().UnixMilli()-time.Now().Add(-time.Hour*2).Add(-time.Minute*1).Add(-time.Second*1).UnixMilli())/float64(1000)/float64(60)))

	float, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(time.Now().UnixMilli()-time.Now().Add(-time.Hour*2).Add(-time.Minute*1).Add(-time.Second*1).UnixMilli())/float64(1000)/float64(60)), 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("+++", float)
}
