package znet

import (
	"fmt"
	"github.com/aceld/zinx/znet"
	"learn/55_zinx/zinx/utils"
	"net"
	"time"

	"learn/55_zinx/zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	msgHandler ziface.IMsgHandle
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
			dealConn := NewConnection(conn, cid, s.msgHandler) // 将connection的连接和handle绑定
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
/*
  创建一个服务器句柄
*/
func NewServer() ziface.IServer {
	//先初始化全局配置文件
	utils.GlobalObject.Reload()

	s := &Server{
		Name:       utils.GlobalObject.Name, //从全局参数获取
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,    //从全局参数获取
		Port:       utils.GlobalObject.TcpPort, //从全局参数获取
		msgHandler: NewMsgHandle(),
	}
	return s
}

// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router succ! ")
}

// PingRouter ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle  pre")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle ----")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle post")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// 现在我们就给用户提供一个自定义的conn处理业务的接口吧，很显然，
// 我们不能把业务处理的方法绑死在type HandFunc func(*net.TCPConn, []byte, int) error这种格式中，
// 我们需要定一些interface{}来让用户填写任意格式的连接处理业务方法。

// HelloZinxRouter Handle
type HelloZinxRouter struct {
	znet.BaseRouter
}

func (h *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}
