package main

import (
	"crypto/tls"
	"time"
)

// 我们将server配置的几个非必须的字段拿出来拼成一个结构体

// 服务器非必填字段
type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

// 服务器配置 需要配置里面的参数
type Server struct {
	Addr   string
	Port   int
	Config *Config
}

// 获取一个默认的服务器配置 所以返回值是server对象的地址
//
func NewDefaultServer(addr string, port int) *Server {
	return &Server{Addr: addr, Port: port}
}

func NewConfig(protocol string) *Config {
	return &Config{
		Protocol: protocol,
	}
}

func NewServer(conf *Config) *Server {
	return &Server{
		Config: conf,
	}
}

// 函数的类型就是func 加上 入参 和 出参

type Option func(s *Server) // 这个变量知道自己指向的函数的签名

// 将函数定义成这样 是我们需要对server对象进行处理

// 函数 函数名 入参 出参 函数体里面的内容

// 上面只是定义了一个函数类型 入参和出参 并没有对函数体里面的内容进行编写

// 也就是说我们将这个函数作为其他函数入参或者出参的时候  就可以使用这个函数的入参

// 函数就是处理函数入参的

// 定义一个函数签名为func(*Server)，函数名为Option 的函数对象

// 其实写个函数就是对某个对象和属性进行操作的一组方法

// 定义一个入参为string类型对象p 出参为Option结构体对象，函数名叫protocol的函数
func Protocol(p int) Option {
	return func(s *Server) {
		s.Port = p
	}
}

// 相当于我将函数protocol对p的处理移到了函数option函数体里面

// 每次我们 传递一个对象 有两种方式
// go语言参数传递是值的传递 形参 是对 实参 值的一份拷贝
//  那为什么入参 不用指针呢

// 函数可作为函数参数进行传递
