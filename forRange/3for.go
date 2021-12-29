package main

import "fmt"

func main() {
	s := []int{0, 1}
	for _, v := range s {
		s = append(s, v)
	}
	fmt.Printf("s=%v\n", s)
}

// 往数组里面追加元素，用for range , for range 会先拷贝一份遍历的数据 , 然后再遍历拷贝的数据，for len 是遍历原数组
