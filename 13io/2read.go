package main

import (
	"fmt"
	"strings"
)

// 1.先创建一个可读的字节流
func main() {
	r := strings.NewReader("Hello, Reader!") // 读出来是一个二进制数据流
	// 2.创建一个字节流的字节数组
	b := make([]byte, r.Size())
	// 3.读取字节流
	n, err := r.Read(b)
	if err != nil {
		fmt.Println("读取失败", err)
	}
	fmt.Println("读取数据长度", n)

	fmt.Println("流中数据", string(b))

}
