package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"net/http"
	"time"
)

type Handle struct{}

func (h *Handle) ServeHTTP(r http.ResponseWriter, request *http.Request) {
	h.Common(r, request)
}

func (h *Handle) Common(r http.ResponseWriter, request *http.Request) {
	hystrix.ConfigureCommand("mycommand", hystrix.CommandConfig{
		Timeout:                int(3 * time.Second),
		MaxConcurrentRequests:  10,
		SleepWindow:            5000, // 五秒后去执行下游服务
		RequestVolumeThreshold: 10,
		ErrorPercentThreshold:  10,
	})
	msg := "success"
	_ = hystrix.Do("mycommand", func() error {
		_, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Printf("请求失败:%v", err)
			return err
		}
		return nil
	}, func(err error) error {
		fmt.Printf("handle  error:%v\n", err)
		msg = err.Error()
		return nil
	})
	r.Write([]byte(msg))
}
func main() {

	log.Fatal(http.ListenAndServe(":8090", &Handle{}))
}
