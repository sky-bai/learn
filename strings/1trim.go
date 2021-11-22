package main

import (
	"fmt"
	"strings"
)

func main() {
	// 去掉字符串首位的字符
	s1 := strings.Trim("abaca", "ba")
	fmt.Println(s1)

}
