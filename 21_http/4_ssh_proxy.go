package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func fiddler() error {

	// 终端输入 ssh -D 127.0.0.1:10086 root@120.77.78.129 进行代理 后面是白名单机器

	// 解析代理 URL
	proxy, err := url.Parse("socks5://127.0.0.1:10086")
	if err != nil {
		log.Println("error: Parse", err)
		return err
	}

	// 创建一个 HTTP 客户端，并设置代理
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	// 创建一个 HTTP GET 请求
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		log.Println("error: NewRequest", err)
		return err
	}

	// 通过客户端发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error: Do", err)
		return err
	}
	// 确保响应的 Body 在函数结束时关闭
	defer resp.Body.Close()

	// 读取响应的 Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error: ReadAll", err)
		return err
	}

	// 打印响应的 Body
	fmt.Println(string(body))

	return nil
}
