package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

func main() {
	tw, err := collection.NewTimingWheel(1*time.Second, 300, func(key, value interface{}) {
		fmt.Println("k", key, "定时器执行的时间time.Now()", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("第一次设置的时间time.now", time.Now().Format("2006-01-02 15:04:05"))
	err = tw.MoveTimer(2, 10*time.Second)
	if err != nil {
		fmt.Println("err1", err)
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Println("第二次设置的时间time.now", time.Now().Format("2006-01-02 15:04:05"))
	err = tw.MoveTimer(2, 10*time.Second)
	if err != nil {
		fmt.Println("err2", err)
		return
	}
	//
	select {}

	// todo 两次move 的处理
	// todo 两次set 的处理

}

// 用户直播时间的统计，与前端页面的心跳检测
