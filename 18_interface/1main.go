package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} = "hello"
	s := i.(string)
	println(s)

	fmt.Println(reflect.TypeOf(i))

	fmt.Println(reflect.ValueOf(i))
}
