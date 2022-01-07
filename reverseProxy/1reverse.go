package main

import (
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
	// "127.0.0.1:2002/dir"
	// "http://127.0.0.1:2003/base/dir"
	proxy := httputil.NewSingleHostReverseProxy(parse) // 能处理res req的handler
	// 转发到目标url后面
	log.Println("Starting httpserver at " + addr)
	http.ListenAndServe(addr, proxy)
}

//  这个反向代理结构体需要一个下游服务器的地址。
///  也就是说为这个handler 是一个反向代理结构体 他也实现了 ServerHttp 方法
//
// 反向代理结构体的作用是 根据传入的url（也就是下游服务器地址） 去构建一个http服务器 当请求这个http服务器的时候 就会把请求转发到 下游的服务器上面

// 反向代理服务器的作用就是将请求转发到设置的服务器上面
