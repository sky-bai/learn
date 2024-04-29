package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(40005009)
	fmt.Println("s", s)
	// 截取后四位
	fmt.Println("---", s[len(s)-4:])
	data := s[len(s)-4:]
	// 转换成int
	i, _ := strconv.Atoi(data)
	fmt.Println("data", i)
}
