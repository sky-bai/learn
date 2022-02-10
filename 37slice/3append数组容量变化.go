package main

import "fmt"

func main() {
	w1 := make([]int, 5)
	fmt.Println(cap(w1))

	w1 = append(w1, 1)
	fmt.Println(cap(w1))

}

// 我这里是直接对数组进行扩容。
// 给一个数组进行append添加元素的时候，会使底层数组进行扩容，扩容的大小为原来的容量的2倍。
