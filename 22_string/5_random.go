package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个0.2 - 0.35之间的随机数
	data := 0.2 + 0.15*rand.Float64()
	// 保留两位小数
	fmt.Printf("%.2f", data)

	// 生成一个0.9 - 1.1之间的随机数
	data = 0.9 + 0.2*rand.Float64()
	// 保留两位小数
	fmt.Printf("%.2f", data)

}
