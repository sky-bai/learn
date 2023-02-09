package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {

	var jsonBlob = []byte(`[     {"Name": "Platypus", "Order": "Monotremata"},     {"Name": "Quoll",    "Order": "Dasyuromorphia"} ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := jsoniter.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	fmt.Printf("%+v", animals[0])
	fmt.Println("111")

	var jsonBlo = []byte(`{"Name": "Platypus", "Order": "Monotremata"}`)
	type Animal1 struct {
		Name  string
		Order string
	}
	var te Animal1
	err = jsoniter.Unmarshal(jsonBlo, &te)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", te)
	aa, err := jsoniter.Marshal(te)
	if err != nil {
		fmt.Println("error:", err)

	}
	fmt.Println("---", aa)
}

// json 解析过后的数据用结构体去接收
// 如果是数组，需要用到切片去接收
// 如果是对象，需要用到结构体去接收
