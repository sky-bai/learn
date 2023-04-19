package main

import (
	"fmt"
	"strings"
)

// 字节是网络传输的基本单位 在内存和磁盘存储信息的单位
// 一个英文占一个字节的大小 一个中文占两个字节的大小
// 字符串转字节数组 字节数组每一个存的是每个字符的编码
func main() {
	bytes := []byte("hello")
	bytes1 := []byte("白")
	fmt.Println(len(bytes1))
	fmt.Println(bytes)
	version := "1.3.4"
	if strings.Compare(version, "1.3.5") < 0 {
		fmt.Println("小于")
	}

}
