package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看机器的核心数
	fmt.Println("---", runtime.NumCPU())
}
