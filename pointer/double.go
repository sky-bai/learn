package main

import "fmt"

func double(x int) {
	x += x
}

func main() {
	a := 3
	double(a) // 函数传参是值传递 double函数的形参是对实参a的一个拷贝
	fmt.Println(a)

}
