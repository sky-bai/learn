package main

import "fmt"

func main() {
	s := []int{5}
	fmt.Println(cap(s)) //1

	s = append(s, 7)
	fmt.Println(cap(s)) //2
	fmt.Println(s)

	s = append(s, 9)
	fmt.Println(cap(s)) //4
	fmt.Println(s)

	x := append(s, 11)
	y := append(s, 12)

	fmt.Println(s, x, y) //[5 7 9] [5 7 9 12] [5 7 9 12]

	//s1:=[5]int{2}
	//fmt.Println(cap(s1)) //5
	//fmt.Println(len(s1)) //5

	//s2 := s1[0:3]
	//fmt.Println(cap(s2)) //5
	//fmt.Println(len(s2)) //3

}

// 对于数组来说 cap和len都是获取该数组实际能存放的元素个数。

// cap 容量就是底层数组实际能存放的元素个数
// len 对于切片来说，长度就是切片的实际长度，而容量则是底层数组实际能存放的元素个数。

// 首先要搞清楚容量和长度的区别：

//容量是指底层数组的大小，长度指可以使用的大小

//容量的用处在哪？在当你用 append扩展长度时，如果新的长度小于容量，不会更换底层数组，否则，go 会新申请一个底层数组，拷贝这边的值过去，把原来的数组丢掉。也就是说，容量的用途是：在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡。

//而长度，则是为了帮助你限制切片可用成员的数量，提供边界查询的。所以用 make 申请好空间后，需要注意不要越界【越 len 】

// slice 切片是一个结构体 里面储存了对原数组的引用，原数组实际存放元素的个数cap 和现在自己本身能够存放的元素个数 len
