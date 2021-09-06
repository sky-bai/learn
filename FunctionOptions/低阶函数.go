package main

import (
	"crypto/tls"
	"time"
)

// 1.定义低阶函数入参 这个参数就是你要操作的对象
// 比如说我们要操作server服务器对象 那么参数就是该对象地址 因为go语言的参数传递是值传递
// 定义低阶函数

type Option1 func(s *Server1) // 创建操作某个对象的函数签名 并没有写具体的方法
// 我定义

// 2.定义高阶函数 需要将高阶函数的入参 传给 低阶函数 让低阶函数处理  因为在上面定义低阶函数的时候就没有实现该函数体
// 所以在下面确定函数体 具体要做的事情

func Protocol1(p string) Option1 {
	// 请注意 函数的入参 和 出参 都是形参
	return func(s *Server1) {
		s.Protocol = p
	}
}

// 所以我需要明确高阶函数的入参 和 低阶函数的入餐

type Server1 struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

func (s Server1) set() {

}

func Timeout1(time time.Duration) Option1 {
	return func(s *Server1) {
		s.Timeout = time
	}
}

// fun(s *server1) 传入一个处理server服务器的函数
func Newserver(add string, port int, option1 ...func(s *Server1)) (*Server, error) {
	// 要学会调用高阶函数和低阶函数 明确高阶函数和低阶函数操作的对象
	return &Server{}, nil
}

// 明确低阶函数操作的对象

// 高阶函数返回低阶函数名   gin的中间件模型应该与这个有关
