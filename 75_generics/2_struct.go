package main

type MyStruct struct {
	Name string
}

// MyStructG 结构体内部的字段不明确 1.定义一个范型类型的结构体 范型类型写在结构体名字后面
type MyStructG[T int | string] struct {
	Name T
}

func main() {

}
