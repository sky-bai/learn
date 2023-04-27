package main

import (
	"container/list"
	"fmt"
	"time"
)

// 在 Go 语言中实现单层时间轮，可以使用以下步骤：
//
// 1. 定义一个结构体，包含一个 tick 值，一个槽数组，以及一个当前槽的索引值。
//
// ```go

type TimeWheel1 struct {
	tick  int64        // 每个tick表示的时间长度
	slots []*list.List // 每个槽存储的所有元素 每个槽是用数组去表示的 每个槽里面存储的是一个链表 后面是这个数组存的是什么 什么类型的数据
	index int          // 当前槽的索引值 时间轮的指针
}

// 一个是它的时间轮的刻度，一个是时间间隔

//```
//
//2. 定义向时间轮中添加元素的方法。该方法将元素添加到指定的槽中。
//
//```go

func (tw *TimeWheel1) AddElement(delay int64, element interface{}) {
	// 多少秒之后执行某个函数 计算任务在哪一个槽位

	// 计算元素应该添加到哪个槽中
	pos := (tw.index + int(delay/tw.tick)) % len(tw.slots)

	// 1.根据延时时间计算出元素应该存放的槽位
	// delay / tick  代表的是多少个tick之后执行

	// 2.表示应该存放在那个槽位上
	// tw.index + int(delay/tw.tick)

	// 3.对24取模
	// 单层时间轮的设计中槽位数量通常是固定的，并且槽位编号是从0开始连续编号的，因此当槽位编号达到槽位数量时，需要对其进行取模操作，以重新回到槽位编号的初始位置，这个操作通常被称为"循环绕回"。
	// 值得注意的是，当计算出的槽位编号大于或等于槽位数量时，也需要进行取模操作，以确保计算出的槽位编号在槽位范围之内。
	// 对于24小时制时间轮，槽位数量通常是24个，因此需要对槽位编号进行24的取模操作，以确保槽位编号的取值在0到23的范围内。这样做可以保证任务都被放入了正确的槽位。
	// (tw.index + int(delay/tw.tick)) % len(tw.slots)

	// 将元素加入到槽中
	tw.slots[pos].PushBack(element)
}

// pos 的计算过程可以分为以下几步：
//
//1. 求出元素的延迟时间 delay 对应的 tick 次数，即 delay/tw.tick。这个值是一个整数，表示元素需要经过多少个 tick 的时间才能被触发。
//
//2. 将上面求得的 tick 次数加上当前时间轮的 index，得到元素应该存放的槽位的绝对位置。
//
//3. 对槽位的绝对位置取模，得到元素在当前时间轮中实际应该存放的槽位位置。
//
//这个计算的意义是将元素按照时间顺序存放到时间轮的槽位中，而且在计算过程中考虑了当前时间轮的索引位置。这样一来，当时间轮的 index 变化时，每个槽位存放的元素也会相应地向下一个槽位移动，以保证能够按照规定的延迟时间触发相应的任务。

//```
//
//3. 定义启动时间轮的方法。该方法按照固定的时间间隔循环执行，每次执行将当前槽中的所有元素取出并执行相应操作。
//
//```go

func (tw *TimeWheel1) Start(interval time.Duration, callback func(interface{})) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		// 取出当前槽中的所有元素，并执行相应操作
		for ele := tw.slots[tw.index].Front(); ele != nil; ele = ele.Next() {
			callback(ele.Value)
		}
		// 将当前槽的索引值加 1，并重新计算
		tw.index = (tw.index + 1) % len(tw.slots)
		// 值得注意的是，当计算出的槽位编号大于或等于槽位数量时，也需要进行取模操作，以确保计算出的槽位编号在槽位范围之内。 会出现什么问题？
	}
}

// ```
//
// 4. 创建时间轮对象并调用相关方法来使用时间轮。
//
// ```go
func main() {
	// 创建一个时间轮对象，tick 值为 1 秒，共有 60 个槽
	tw := &TimeWheel1{
		tick:  1,
		slots: make([]*list.List, 12),
		index: 0,
	}
	for i := 0; i < len(tw.slots); i++ {
		tw.slots[i] = list.New()
	}

	// 将元素添加到时间轮中
	tw.AddElement(1, "task 1")
	tw.AddElement(3, "task 2")
	tw.AddElement(5, "task 3")
	tw.AddElement(7, "task 4")

	// 启动时间轮，并在每个时间间隔中执行相应操作
	tw.Start(time.Second, func(elem interface{}) {
		// 取出元素，并执行相应任务
		task := elem.(string)
		fmt.Println("execute task:", task)
	})
}

//```
//
//以上是使用 Go 语言实现单层时间轮的基本步骤。需要注意的是，在实际生产环境中，可能需要对时间轮的性能进行优化，并添加相应的容错机制以及异常处理逻辑等。
