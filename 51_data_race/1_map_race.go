package main

import (
	"os"
	"os/signal"
	"syscall"
)

var qqMapRunMap map[int]struct{}

func main() {

	for i := 0; i < 10; i++ {
		go NewQQMapDeviceConsumer(i, "platform")
	}

	// 等待中断信号
	quit := make(chan os.Signal)

	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
func init() {
	qqMapRunMap = make(map[int]struct{}, 3000) // 每天的数据
}

func NewQQMapDeviceConsumer(imei int, platform string) {

	if _, ok := qqMapRunMap[imei]; !ok {
		qqMapRunMap[imei] = struct{}{}

	}
}
