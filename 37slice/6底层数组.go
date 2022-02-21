package main

import "fmt"

func main() {
	a := []int{3, 2, 3, 4, 5}

	fmt.Printf("%+v\n", a)

	ap(a)
	fmt.Printf("%+v\n", a)

	app(a)
	fmt.Printf("%+v\n", a)

}

func ap(a []int) {
	a = append(a, 6)
	fmt.Printf("3333%+v\n", a)

}

func app(a []int) {
	a[0] = 1
	fmt.Printf("1111%+v\n", a)

}

// 不会影响原切片长度
// 但是修改值会影响原切片值
