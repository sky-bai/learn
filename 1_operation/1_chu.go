package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println(5 % 2)           // 求出余数
	fmt.Println(float64(13) / 3) // 求出商
	fmt.Printf("Size in Megabytes: %.2f MB\n", float64(13)/3)
	formattedValue := strconv.FormatFloat(float64(13)/3, 'f', 2, 64)
	fmt.Println("保留两位小数 string类型:", formattedValue)

	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(22)/3), 64)
	fmt.Println("保留两位小数 float64类型:", num)

	fmt.Println(2 % 10)

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

	fmt.Println("---", 1/2) // 求出商

}
