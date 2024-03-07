package main

import (
	"fmt"
	"strings"
)

func main() {
	index := strings.Index("Hello World", "World")
	// 返回这个字符串的第一次出现的下标位置
	fmt.Println(index)

	host := "www.baidu.com:8080"
	//host = host[0:strings.Index(host, ":")]
	//fmt.Println(host)

	host = strings.Split(host, ":")[0]
	fmt.Println(host)

	url := "/index.html"
	url = strings.TrimPrefix(url, "/")
	fmt.Println(url)

	path := "/index.html"
	path = strings.TrimPrefix(path, "/")
	fmt.Println(path)

}
