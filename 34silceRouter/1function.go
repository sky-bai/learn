package main

//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//)
//
//type HandlerFunc func(w http.ResponseWriter, r *http.Request)
//
//func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	f(w, r)
//}
//
//func main() {
//	hf := HandlerFunc(HelloHandler)
//
//	// 构建返回值
//	resp := httptest.NewRecorder()
//	req := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("test")))
//
//	hf.ServeHTTP(resp, req)
//	bts,_:=ioutil.ReadAll(resp.Body)
//	fmt.Println(string(bts))
//}
//
//func HelloHandler(res http.ResponseWriter, req *http.Request) {
//	res.Write([]byte("hello")) // write
//}

// 函数实现的方法可以调用这个函数本身
// 请记住 type 后面是类型 是这个函数的类型 通过后面传入这个类型的具体函数获得这个类型的变量 然后再去使用这个类型已有的方法
// 类型加上函数名 获得全这个类型的变量 只不过这个变量是一个函数
// 定义这个类型所要做的事情

// ServerHTTP是http handler接口里面的方法

// 构建一个类型 这个类型是一个函数 该函数的签名是能够处理请求和响应的函数
// 指定这个函数能够处理http请求的方法 然后构建请求和恢复

// 创建一个类型为函数的变量的时候 通过该类型后面加一个实现了这个类型签名的函数去获得这个变量
