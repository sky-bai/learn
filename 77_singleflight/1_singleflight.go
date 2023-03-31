package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var count int32

// 模拟数据库调用
func getArticle(id int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	atomic.AddInt32(&count, 1)
	fmt.Println("count: ", count)
	time.Sleep(time.Duration(count) * time.Millisecond)

	return fmt.Sprintf("article: %d", id), nil
}

func singleFlightGetArticle(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id) // 在Do方法里面执行实际业务操作
	})

	return v.(string), err
}

func main() {
	// 使用一个定时器
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ { // 同时发起1000次请求
		wg.Add(1)
		go func() {
			res, _ := singleFlightGetArticle(sg, 1)
			//res, _ := getArticle(1)
			if res != "article: 1" {
				panic("err")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}

// 1.什么时候用do
// 2.什么时候用doChan
