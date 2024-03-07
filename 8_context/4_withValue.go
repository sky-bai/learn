package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1.创建父节点
	parentCtx := context.Background()
	// 2.加入超时时间
	deadLine, _ := context.WithDeadline(parentCtx, time.Now().Add(55*time.Second))
	context.WithCancel(deadLine)
	// 3.加入值
	newCtx := context.WithValue(deadLine, "startTime", time.Now())
	fmt.Println(newCtx.Value("startTime"))
	for {
		time.Sleep(1 * time.Second)
		t, _ := newCtx.Deadline()
		fmt.Println("-----", t)
	}
}
