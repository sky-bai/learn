package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("D:/go.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//  f是一个*file文件类型 它实现了io.Reader这个接口
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
