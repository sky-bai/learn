package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var addr = "127.0.0.1:2002"

func main() {
	rs1 := "http://127.0.0.1:2003/base"
	parse, err := url.Parse(rs1) //将string解析成*URL格式
	if err != nil {
		log.Println(err)
	}
	fmt.Println(parse)
	// "127.0.0.1:2002/dir"
	// "http://127.0.0.1:2003/base/dir"
	proxy := httputil.NewSingleHostReverseProxy(parse) // 能处理res req的handler
	// 转发到目标url后面
	http.ListenAndServe(addr, proxy)
}
