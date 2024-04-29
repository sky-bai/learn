package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	paramMap := make(map[string]interface{})
	paramMap["token"] = "SC-7z7Nk9q9me5V5uDfqLrKIbwBu6QdhXRg"
	paramMap["signature"] = "9c0f96468da5bf1222b66978b0d2ff5d"

	value, err := jsoniter.Marshal(paramMap)
	if err != nil {

		return
	}
	fmt.Println("111", paramMap)

	paramMap1 := make(map[string]interface{})

	err = jsoniter.Unmarshal(value, &paramMap1)
	if err != nil {
		return

	}

	fmt.Println("222", paramMap1)

	switchCase()
}

func switchCase() {

	i := 3
	switch i {
	case 1, 2:
		fmt.Println(" --- 1")
	case 3:
		fmt.Println(" --- 3")
	}
}
