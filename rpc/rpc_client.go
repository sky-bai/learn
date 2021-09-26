package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 算数运算请求结构体
type ArithRequest1 struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse1 struct {
	Pro int // 乘积
	Que int // 商
	Rem int // 余数
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095") // 1.连接服务
	if err != nil {
		log.Fatalln("dailing error:", err)
	}
	req := ArithRequest1{9, 2}
	var res ArithResponse1
	err = conn.Call("Arith.Multipy", req, &res) // 2.调用服务端的服务
	if err != nil {
		log.Fatalln("Arith error:", err)
	}
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("Arith error:", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Que, res.Rem)
}
