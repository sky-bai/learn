package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "jhfhDfXE5k<<<#1:803277203614964:1:*,00000B5E,STATUS,29,59339,3,1,1,0,0,0,0,0,0,0,0.000000,-89,1,-1,-1#"
	arr := strings.Split(msg, ",")
	fmt.Println(len(arr))
}
