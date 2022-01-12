package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	s := a[1:4]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。 注意底层数组
}
