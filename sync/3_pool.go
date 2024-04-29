package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
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

func main() {
	data := []string{"hello", "sdfaf"}
	dataStr, _ := json.Marshal(data)
	fmt.Println(string(dataStr))

	fmt.Println(time.Now().Format("01-02 15:04:05"))
}
