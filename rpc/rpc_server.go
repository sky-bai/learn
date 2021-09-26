package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Que int // 商
	Rem int // 余数
}

// 乘法运算方法
func (a *Arith) Multipy(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (a *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Que = req.A / req.B // 商
	res.Rem = req.A % req.B // 余数
	return nil
}

func main() {
	rpc.Register(new(Arith)) // 1.注册rpc服务
	rpc.HandleHTTP()         // 2.采用http协议作为rpc载体

	listen, err := net.Listen("tcp", "127.0.0.1:8095") // 3.启动服务
	if err != nil {
		log.Fatalln("fatal error", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection")
	http.Serve(listen, nil)

}
