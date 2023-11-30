package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

func main() {
	// 获取当前时间到当天最后一秒的间隔
	fmt.Println(now.EndOfDay().Sub(time.Now()).Seconds())
}
