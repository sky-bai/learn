package main

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println(s1[1:3])
	for index := range s1 {
		fmt.Println("for range 遍历切片后 只获取一个值，该值是索引下标", index)
		s1[index] *= 2
	}
	fmt.Printf("%v\n", s1)
}
