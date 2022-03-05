package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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
	proxy := NewSingleHostReverseProxy(parse) // 能处理res req的handler
	// 转发到目标url后面
	log.Println("Starting httpserver at " + addr)
	http.ListenAndServe(addr, proxy)
}

// NewSingleHostReverseProxy 提供一个修改请求url反向代理的结构体 需要提供
func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery        // ？后面的查询参数 a=111&b=456
	director := func(req *http.Request) { // 将原请求修改为目标请求
		req.URL.Scheme = target.Scheme // scheme http
		req.URL.Host = target.Host
		req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL) // 确定新的url
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	// 修改回复
	modifyFunc := func(response *http.Response) error {
		if response.StatusCode != 200 {

			// 拿到返回内容 响应的内容在http的body里面
			oldPayload, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}

			// 1.修改body 和 body对应的contentLength

			// 在后面追加内容
			newPayLoad := []byte("hello" + string(oldPayload))
			// 将内容回写到回复里面 这里需要一个bytes转io.ReadCloser的方法
			response.Body = io.NopCloser(strings.NewReader(string(newPayLoad)))
			// 一个具体的类型实现了某一接口的方法，那么该类型就实现了该接口，该类型变量就是该接口变量
			// response.Body 是一个readCloser接口 那么我就要让我的内容实现readCloser接口里面的方法 把该接口变量转换成readCloser接口变量

			// 内容长度是按照body的长度来的
			response.ContentLength = int64(len(newPayLoad))

			// 我们要让客户端知道他读到响应的大小，所以我们在header里面设置
			response.Header.Set("Content-Length", fmt.Sprint(len(newPayLoad)))
		}
		return nil
	}

	// 官方提供的new方式只是一个对某一内容的实现，reverserProxy是实现功能的结构体，我们可以通过New方法进行设置我们想要的ReverseProxy结构体去实现我们想要的功能
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc}
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
