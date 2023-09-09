package main

import "fmt"

// 1.定义一个范型类型的map
type myMap[key int | string, value float32 | float64] map[key]value

func main() {
	var a myMap[int, float32] = map[int]float32{1: 1.1}
	fmt.Println(a)
}
