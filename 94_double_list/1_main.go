package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.List{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	fmt.Println("l.Len()", l.Len())

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("e.Value", e.Value)
	}
}

// 这是使用golang标准库container/list中的双向链表数据结构l的一个遍历写法。
// 其中，l.Front()返回链表l中的第一个元素；e := l.Front()将e初始化为链表中第一个元素；e != nil表示如果e不为nil，则循环继续；e = e.Next()将e更新为链表中下一个元素，最后达到遍历整个链表的效果。
//
// 在每次循环中，输出e节点上的值e.Value，即完成遍历操作。
