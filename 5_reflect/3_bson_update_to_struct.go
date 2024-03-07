package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type Vehicle struct {
	Brand string `json:"brand" bson:"brand"`
	Other string `json:"other" bson:"other"`
}
type te struct {
	Name string `json:"name"`
}

func main() {

	var v Vehicle
	v.Brand = "宝马"
	v.Other = "其他"

	updateInfo := bson.M{
		"brand": "奥迪",
	}

	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("bson")
		if value, ok := updateInfo[jsonTag]; ok {
			reflect.ValueOf(&v).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
		}
	}

	fmt.Println(v)

	connections := make([]string, 0, 5000)
	for i := 0; i < 580; i++ {
		connections = append(connections, "1")
	}
	for i := 0; i < 4000; i++ {
		connections = append(connections, "123")
	}

	testCount := 0
	testIdArr := make([]string, 0, 500)
	for _, conn := range connections {

		// 如果是测试链接，记录tcp日志
		if conn == "1" {
			testCount++
			testIdArr = append(testIdArr, conn)
		}

		if testCount >= 500 {
			fmt.Println("for len:", len(testIdArr))
			testCount = 0
			testIdArr = make([]string, 0, 500)

		}
	}

	if testCount > 0 {
		fmt.Println("len:", len(testIdArr))
		fmt.Println(testIdArr)
	}

}
