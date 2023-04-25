package main

import "container/list"

// 在 Go 语言中实现单层时间轮，可以使用以下步骤：
//
//1. 定义一个结构体，包含一个 tick 值，一个槽数组，以及一个当前槽的索引值。
//
//```go
//type TimeWheel struct {
//    tick int64                // 每个tick表示的时间长度
//    slots []*list.List        // 每个槽存储的所有元素
//    index int                 // 当前槽的索引值
//}
//```
//
//2. 定义向时间轮中添加元素的方法。该方法将元素添加到指定的槽中。
//
//```go
//func (tw *TimeWheel) AddElement(delay int64, element interface{}) {
//    // 计算元素应该添加到哪个槽中
//    pos := (tw.index + int(delay/tw.tick)) % len(tw.slots)
//    // 将元素加入到槽中
//    tw.slots[pos].PushBack(element)
//}
//```
//
//3. 定义启动时间轮的方法。该方法按照固定的时间间隔循环执行，每次执行将当前槽中的所有元素取出并执行相应操作。
//
//```go
//func (tw *TimeWheel) Start(interval time.Duration, callback func(interface{})) {
//    ticker := time.NewTicker(interval)
//    for range ticker.C {
//        // 取出当前槽中的所有元素，并执行相应操作
//        for ele := tw.slots[tw.index].Front(); ele != nil; ele = ele.Next() {
//            callback(ele.Value)
//        }
//        // 将当前槽的索引值加 1，并重新计算
//        tw.index = (tw.index + 1) % len(tw.slots)
//    }
//}
//```
//
//4. 创建时间轮对象并调用相关方法来使用时间轮。
//
//```go
//func main() {
//    // 创建一个时间轮对象，tick 值为 1 秒，共有 60 个槽
//    tw := &TimeWheel{
//        tick: 1,
//        slots: make([]*list.List, 60),
//        index: 0,
//    }
//    for i := 0; i < len(tw.slots); i++ {
//        tw.slots[i] = list.New()
//    }
//
//    // 将元素添加到时间轮中
//    tw.AddElement(10, "task 1")
//    tw.AddElement(20, "task 2")
//    tw.AddElement(30, "task 3")
//    tw.AddElement(40, "task 4")
//
//    // 启动时间轮，并在每个时间间隔中执行相应操作
//    tw.Start(time.Second, func(elem interface{}) {
//        // 取出元素，并执行相应任务
//        task := elem.(string)
//        fmt.Println("execute task:", task)
//    })
//}
//```
//
//以上是使用 Go 语言实现单层时间轮的基本步骤。需要注意的是，在实际生产环境中，可能需要对时间轮的性能进行优化，并添加相应的容错机制以及异常处理逻辑等。

type TimeWheel struct {
	tick  int64        // 每个tick表示的时间长度
	slots []*list.List // 每个槽存储的所有元素
	index int          // 当前槽的索引值
}
