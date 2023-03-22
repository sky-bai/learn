package main

// 从多个网络地址获得信息，如果有一个出错或超时，则取消所有未完成的操作

import (
	"context"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type query struct {
	url  string // 查询地址
	resp string // 正常返回的信息
	err  error  // 错误消息，如果有的话
}

func search(ctx context.Context, q query, resultChan chan<- query) func() error {
	return func() error {
		// 产生新的timeout的context，超时1秒，事实上超时掉/slow请求
		reqCtx, cancel := context.WithTimeout(ctx, time.Second)
		// 永远调用cancel()
		defer cancel()

		req, err := http.NewRequest(http.MethodGet, q.url, nil)
		// go的错误处理让人不快
		if err != nil {
			q.err = err
			resultChan <- q
			return q.err
		}
		// 用新的context创建新的request
		resp, err := http.DefaultClient.Do(req.WithContext(reqCtx))
		// zero-overhead deterministic exceptions are needed
		if err != nil {
			q.err = err
			resultChan <- q
			return q.err
		}

		body, err := ioutil.ReadAll(resp.Body)
		// 让人厌倦
		if err != nil {
			q.err = err
			resultChan <- q
			return q.err
		}
		_ = resp.Body.Close()

		q.resp = string(body)
		resultChan <- q
		return nil
		// 不论是否出错，都讲查询结果送入resultChan
	}
}

func fetch(ctx context.Context, queries []query, resultChan chan<- query) error {
	// 本函数结束时，所有的结果都应该已经进入此channel，所以关闭
	defer close(resultChan)

	// 产生新的context
	eg, egCtx := errgroup.WithContext(ctx)
	for _, q := range queries {
		// errGroup.Go接受func() error
		// 启动所有的go routine，并传入context
		eg.Go(search(egCtx, q, resultChan))
	}

	return eg.Wait()
}

func main() {
	var queries = []query{
		{url: "http://localhost:45678/fast"}, // "立即"返回信息
		{url: "http://localhost:45678/slow"}, // 等待两秒后返回，因为超时时间设为一秒，实际上此条超时
	}
	// 结果都进入这个channel，长度和请求的个数一样，所以不阻塞
	var resultsChan = make(chan query, len(queries))
	// 以下是实际请求和结果输出
	if err := fetch(context.Background(), queries, resultsChan); err != nil {
		log.Printf("get error %v", err)
	} else {
		log.Print("everything is fine")
	}
	for r := range resultsChan {
		if r.err != nil {
			log.Printf("%s error: %v", r.url, r.err)
			continue
		}
		log.Printf("%s result: %s", r.url, r.resp)
	}
}
