package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

// ImeiBlackList 保存imei黑名单
var ImeiBlackList = &ImeiBlackListStruct{}

type ImeiBlackListStruct struct {
	QqMapBlackList *sync.Map
}

func NewBlackList() *ImeiBlackListStruct {
	return &ImeiBlackListStruct{
		QqMapBlackList: &sync.Map{},
	}
}

func (i *ImeiBlackListStruct) Start() {
	i.QqMapBlackList.Store("123456", "12312")
	value, ok := i.QqMapBlackList.Load("123456")
	fmt.Println("第一次", value, ok)

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	rdb.HSet(context.Background(), "test_hgetall", "key1", "value1")
	result, err := rdb.HGetAll(context.Background(), "test_hgetall").Result()
	if err != nil {
		fmt.Println(err)
		return

	}
	var syncMap sync.Map
	for k, _ := range result {
		syncMap.Store(k, struct{}{})
	}
	i.QqMapBlackList = &syncMap

	value1, ok1 := i.QqMapBlackList.Load("key1")
	if ok1 {
		fmt.Println("ok", value1)
	} else {
		fmt.Println("not ok")
	}

	value2, ok2 := i.QqMapBlackList.Load("123456")
	fmt.Println("第二次", value2, ok2)

}

func main() {
	NewBlackList().Start()
}
