package main

import "fmt"

type AnySlice[T int | float64 | string] []T

// 定义一个可以有int float64 string类型的切片
// int|float32|float64 这部分被称为类型约束(Type constraint)，中间的 | 的意思是告诉编译器，类型形参 T 只可以接收 int 或 float32 或 float64 这三种类型的实参
// 中括号里的 T int|float32|float64 这一整串因为定义了所有的类型形参(在这个例子里只有一个类型形参T），所以我们称其为 类型形参列表(type parameter list)

// IntSlice 1,定义一个任意类型的切片
type IntSlice[T any] []T

func main() {
	// 1.实例化
	var int IntSlice[int] = []int{1, 2, 3}
	fmt.Println(int)

}
