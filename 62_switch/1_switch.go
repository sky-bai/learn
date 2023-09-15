package main

import (
	"github.com/tal-tech/go-zero/core/timex"
	"time"
)

func main() {
	// 创建一个时间轮实例，精度为1秒，共有60个槽
	timingWheel := timex.NewTimingWheel(time.Second, 60)
	timex.Time()
}
