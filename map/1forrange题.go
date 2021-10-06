package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, value := range slice {
		fmt.Println(key, &value)
		// value是一个固定不变的地址 也就是说m[]后面只有一个地址 保存者最后一个值
		m[key] = &value
	}
	fmt.Println(*m[2])
}
