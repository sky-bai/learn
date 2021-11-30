package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.reader stream to be read\n")
	// 实现了reader接口的结构体 表示可以从里面读数据
	// 实现了writer接口的结构体 表示可以往里面写数据
	if _, err := io.Copy(os.Stdout, r); err != nil {
		panic(err)
	}

}

// write 方法将字节流写入底层数据流
