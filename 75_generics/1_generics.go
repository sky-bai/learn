package main

import "fmt"

type AnySlice[T int | float64 | string] []T

// 定义一个可以有int float64 string类型的切片
// int|float32|float64 这部分被称为类型约束(Type constraint)，中间的 | 的意思是告诉编译器，类型形参 T 只可以接收 int 或 float32 或 float64 这三种类型的实参
// 中括号里的 T int|float32|float64 这一整串因为定义了所有的类型形参(在这个例子里只有一个类型形参T），所以我们称其为 类型形参列表(type parameter list)

// IntSlice 1. 定义一个任意类型的切片
type IntSlice []int
type IntSliceG[T any] []T

// MyStruct 2.定义一个范型类型的结构体 范型类型写在结构体名字后面
type MyStruct struct {
	Name string
}
type MyStructG[T int | string] struct {
	Name T
}

// 部标

// MyInterface 3.定义一个范型接口
type MyInterface interface{}
type MyInterfaceG[T any] interface {
	Write(v T)
}

// IPrintData 4.一个泛型接口(对于泛型接口在后半部分会详细讲解）
type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

// MyChan 4. 一个泛型通道，可用类型实参 int 或 string 实例化
type MyChan chan int
type MyChanG[T int | string] chan T

func main() {
	// 1.实例化切片
	var int IntSliceG[int] = []int{1, 2, 3}
	fmt.Println(int)

	// 2.实例化结构体
	var myStruct MyStruct
	myStruct.Name = "hello"

	var myStructG MyStructG[string]
	myStructG.Name = "hello"

	// 可是对于
}
