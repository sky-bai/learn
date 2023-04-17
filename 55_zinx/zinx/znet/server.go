package znet

import (
	"fmt"
	"github.com/aceld/zinx/znet"
	"net"
	"time"

	"learn/55_zinx/zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	//当前Server由用户绑定的回调router,也就是Server注册的链接对应的处理业务
	Router ziface.IRouter
}

// Start 启动服务器
func (s *Server) Start() {
	// TODO
	go func() {
		// 1.获取本机服务器地址然后构造对象
		add, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			panic(err)
		}
		// 2.启动服务器开始监听
		listener, err := net.ListenTCP(s.IPVersion, add)
		if err != nil {
			panic(err)
			return
		}
		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0
		// 3.阻塞等待客户端连接，处理客户端连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			//3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(conn, cid, s.Router) // 将connection的连接和handle绑定
			cid++

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}

}

// NewServer 创建一个服务器句柄
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
		Router:    nil,
	}

	return s
}

// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router

	fmt.Println("Add Router succ! ")
}

// PingRouter ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

// 传入一个9个字节长的数据，分多次put （对应于TCP中的分包的情况）
// 传入一个3个字节的数据和一个6个字节的数据，一次put（对应于TCP中的粘包的情况）
// 大数据处理测试 (20MB)

// 缓存区长度，默认8k 使得缓存区的长度可以自定义

// 现在我们就给用户提供一个自定义的conn处理业务的接口吧，很显然，
// 我们不能把业务处理的方法绑死在type HandFunc func(*net.TCPConn, []byte, int) error这种格式中，
// 我们需要定一些interface{}来让用户填写任意格式的连接处理业务方法。
