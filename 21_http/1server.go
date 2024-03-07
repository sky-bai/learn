package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 函数作为一种数据类型

	//hf := HandlerFunc(HelloHandler)
	//hf.ServeHttp()
}

// HandlerFunc 这是一个函数类型 这是个类型 不是变量 通过传入与这个函数签名相同的函数，就可以得到这个函数类型的变量
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// ServeHttp 这里是给这个类型定义一个方法 只要是这个类型的函数就有下面这个方法
func (f HandlerFunc) ServeHttp(w http.ResponseWriter, r *http.Request) {
	// 这里回调了本身这个函数类型的方法
	f(w, r)
}

func HelloHandler(res http.ResponseWriter, r *http.Request) {

}
