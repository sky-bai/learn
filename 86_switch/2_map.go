package main

import (
	"fmt"
	"time"
)

var TestMap map[string]string

func init() { TestMap = make(map[string]string, 1) }
func main() {
	for i := 0; i < 100; i++ {
		go Write("aaa")
		go Read("aaa")
		go Write("bbb")
		go Read("bbb")
	}
	time.Sleep(5 * time.Second)
}
func Read(key string)  { fmt.Println(TestMap[key]) }
func Write(key string) { TestMap[key] = key }
