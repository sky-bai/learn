package main

import timewheel "github.com/zeromicro/go-zero/core/collection"

func main() {
	tw, _ := timewheel.NewTimingWheel(1, 20, func(key, value interface{}) {
		println("hello")
	})

	tw.SetTimer(1, tw, 1)

	select {}
}

// 服务断掉，那我后台岂不是还是直播着的
