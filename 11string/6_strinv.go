package main

import (
	"fmt"
	"strconv"
)

func main() {

	// uint64转string
	var i uint64 = 123456789
	fmt.Println(strconv.FormatUint(i, 10))

	// string 转 uint64
	var s string = "123456789"
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
