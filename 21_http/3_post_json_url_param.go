package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	// 创建一个URL对象
	reqURL, _ := url.Parse("https://xxx.xxx.xx.xxx:8009/login")

	// 创建一个URL参数对象
	params := url.Values{}
	params.Add("sid", "sid")

	// 将参数编码并追加到URL的查询字符串中
	reqURL.RawQuery = params.Encode()

	fmt.Println(reqURL.String())

	// 创建一个POST请求
	req, _ := http.NewRequest("POST", reqURL.String(), nil)

	// 发送请求
	client := &http.Client{}
	resp, _ := client.Do(req)

	// 处理响应
	// ...
	fmt.Printf("%+v\n", resp)
}
