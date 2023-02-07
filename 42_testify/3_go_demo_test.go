package main

import "github.com/stretchr/testify/mock"

type MyTestObject struct {
	// add a Mock object instance
	mock.Mock

	// other fields go here as normal
}

func (o *MyTestObject) SavePersonDetails(firstname, lastname string, age int) (int, error) {
	args := o.Called(firstname, lastname, age) // 获取mock对象参数
	return args.Int(0), args.Error(1)
}
