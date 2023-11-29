package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

// 直接在项目里面搜sync.Once

func main() {
	fmt.Println(now.EndOfDay().Add(time.Hour * -4).Sub(time.Now()).Seconds())
	fmt.Println(now.EndOfDay().Add(time.Hour * -4).Sub(time.Now()).Seconds())

}
