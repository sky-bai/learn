package main

import (
	"fmt"
	"math"
)

func main() {
	timestamp1 := int64(1635938400000) // 示例时间戳1，表示2021年11月3日 12:00:00.000
	timestamp2 := int64(1635938401000) // 示例时间戳2，表示2021年11月3日 12:00:05.000

	// 计算时间戳之间的差值
	timeDifference := timestamp1 - timestamp2
	fmt.Println("时间戳1和时间戳2的差值为：", timeDifference)

	fmt.Println(math.Abs(float64(timeDifference)))

	// 判断是否在一秒之内
	if math.Abs(float64(timeDifference)) <= 1000 {
		fmt.Println("时间戳1和时间戳2在一秒之内")
	} else {
		fmt.Println("时间戳1和时间戳2不在一秒之内")
	}
}
