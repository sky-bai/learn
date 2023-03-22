package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前路径
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	fileName := dir + "/test.txt"
	create, err := os.Create(fileName)
	if err != nil {
		fmt.Println("---", err)
		return
	}
	fmt.Println("-=1")
	_, err = create.WriteString("hello world1")
	if err != nil {
		fmt.Println("---", err)
		return
	}
	defer create.Close()

}
