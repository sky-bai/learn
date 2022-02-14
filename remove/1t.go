package main

import (
	"fmt"
	"strings"
)

func main() {

	taskArrIds := "[1, 2, 3, 4, 5]"
	// 转变成数组
	//taskArrIds = strings.Replace(taskArrIds, "[", "", -1)
	fmt.Println(taskArrIds)
	str1 := strings.ReplaceAll(taskArrIds, "[", "")
	str2 := strings.ReplaceAll(str1, "]", "")
	fmt.Println(str2)

	idArrayStr := strings.Split(str2, ",")
	for _, V := range idArrayStr {
		fmt.Println(V)
	}

}
