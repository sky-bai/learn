package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string `json:"name"`
}

func main() {
	m1 := MyStruct{
		Name: "白",
	}
	myType := reflect.TypeOf(m1) // 返回m1的动态类型
	fmt.Println("结构体的类型", myType)

	field1 := myType.Field(0)
	fmt.Println("结构体的第一个字段", field1)

	tag := field1.Tag.Get("json")
	fmt.Println("结构体字段的标签", tag)

	myMap := make(map[string]string, 10)
	myMap["name"] = "白"
	typeOf := reflect.TypeOf(myMap)
	fmt.Println("typeof", typeOf)

	value := reflect.ValueOf(myMap)
	fmt.Println("valueOf", value)
}
