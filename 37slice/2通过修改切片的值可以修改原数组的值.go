package main

import "fmt"

func main() {

	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100
	fmt.Println(foo)

	bar := foo[1:4]
	bar[1] = 99

	fmt.Println(foo)
}

// 切片与数组数据共享
// s[low : high : max]
//从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low
