package main

import (
	"encoding/json"
	"sync"
)

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
