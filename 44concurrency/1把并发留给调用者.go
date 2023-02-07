package main

import (
	"context"
	"fmt"
	"time"
)

// 搞清楚你起的goroutine 什么时候能够结束
// 你有没有一个手段能够控制他结束

// 往channel写的主动方 决定何时关闭 channel

// 什么时候goroutine 能够结束
// 以及有没有一个办法能够让goroutine 能够结束
// 往在一个做事情的函数参数里面放入一个 chan 信号 通知该任务结束

type Tracker struct {
	ch   chan string   // 是worker工作模型 通过goroutine来消费channel 里面的数据
	stop chan struct{} // 有一种机制能让消费ch的goroutine进行暂停 close(t.ch)
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}
func (t *Tracker) Shutdown(ctx context.Context) {
	// 能写这个channel的owner 才能关闭channel
	close(t.ch) // 这里暂停之后 就没法往ch里面发送数据了 Run方法将消费完ch剩下的数据 通知goroutine进行暂停 这里有ctx控制超时，如果track的任务在规定时间没有完成也会退出
	select {
	case <-ctx.Done():
	case <-t.ch:
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	// 消费channel里的数据
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()

	}
}

func main() {
	tr := NewTracker()
	go tr.Run() // 将事情的并发交给控制者
	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}

// feature1

// main
