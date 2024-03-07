package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	getwd, err := os.Getwd() // 获取当前工作目录
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getwd)
	os.Stat(getwd) // 获取文件信息
	filepathString := filepath.Dir(getwd)
	fmt.Println("filepathString", filepathString)
}
