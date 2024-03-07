package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 42
	ptr := &num

	// 获取指针所指向的变量的类型
	ptrType := reflect.TypeOf(ptr).Elem()

	// 获取指针所指向的变量的类型的名称
	ptrTypeName := ptrType.Name()

	fmt.Println("ptrTypeName:", ptrTypeName) // 输出：int
}

//Worker里面的WorkerPool chan chan Job是什么意思？
//type Job struct {
//   Payload Payload
//}
//
//var JobQueue chan Job
//
//type Worker struct {
//   WorkerPool  chan chan Job
//   JobChannel  chan Job
//   quit        chan bool
//}
// 在给定的代码中，有一个 `Worker` 结构体和一个 `Job` 结构体，用于实现并发的工作池模式。
//
//`WorkerPool` 是一个管道（channel），其类型是 `chan chan Job`。它用于存放可用的工作通道。每个工作通道 `chan Job` 用于接收任务，并将任务分配给可用的工作者（Worker）进行处理。
//
//`JobChannel` 是一个管道（channel），其类型是 `chan Job`。它用于接收要执行的任务。
//
//`quit` 是一个管道（channel），其类型是 `chan bool`。它用于向工作者（Worker）发送信号，告诉它们停止工作并退出。
//
//整体来说，`Worker` 结构体中的 `WorkerPool` 管道用于存放可用的工作通道，`JobChannel` 管道用于接收要执行的任务，`quit` 管道用于发送停止信号。这样组合起来，可以实现一种工作池的机制，其中多个工作者（Worker）从工作池中获取任务并进行处理。
