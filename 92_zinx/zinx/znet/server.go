package znet

import (
	"fmt"
	"io"
	"learn/92_zinx/zinx/utils"
	"learn/92_zinx/zinx/ziface"

	"net"
	"time"
)

// Server iServer 接口实现，定义一个Server服务类
type Server struct {
	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int

	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	msgHandler ziface.IMsgHandle

	//当前Server的链接管理器
	ConnMgr ziface.IConnManager

	//该Server的连接创建时Hook函数
	OnConnStart func(conn ziface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn ziface.IConnection)
}

//============== 实现 ziface.IServer 里的全部接口方法 ========

// Start 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	//开启一个go去做服务端Linster业务
	go func() {

		//0 启动worker工作池机制
		s.msgHandler.StartWorkerPool()

		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		//已经监听成功
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")

		//TODO 1_server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0

		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				if err == io.EOF {
					fmt.Println("listenner is closed")
					break
				}
				fmt.Println("Accept err ", err)
				continue
			}

			//3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			//=============
			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			//3.3 处理该新连接请求的业务方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConntion(s, conn, cid, s.msgHandler)
			cid++

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(1 * time.Second)
	}
}

/*
NewServer 创建一个服务器句柄
*/
func NewServer() ziface.IServer {
	//先初始化全局配置文件
	utils.GlobalObject.Reload()

	s := &Server{
		Name:       utils.GlobalObject.Name, //从全局参数获取
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,    //从全局参数获取
		Port:       utils.GlobalObject.TcpPort, //从全局参数获取
		msgHandler: NewMsgHandle(),             //msgHandler 初始化
		ConnMgr:    NewConnManager(),           //创建ConnManager
	}
	return s
}

// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}

// GetConnMgr 得到链接管理
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

// SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

// 要应对海量请求，Go语言的TCP服务器可以采用以下策略：
//
//1. 使用goroutine和channel：Go语言中的goroutine和channel可以帮助服务器支持并发处理请求。每个连接可以使用goroutine进行处理，goroutine可以并发处理多个连接，而channel则可以用于协调不同goroutine之间的通信。
//
//2. 设置优化参数：Go语言的服务器可以设置一些优化参数，如TCP连接超时时间、TCP发送和接收缓冲区大小等。通过合理地设置这些参数，可以避免服务器由于连接超时或缓冲区不足而导致性能下降。
// 缓存区一般设置为多大昵 这个tcp连接超时时间是客户端连接超时时间还是服务端连接超时时间昵
//3. 使用连接池：连接池是一种重用连接的技术。对于TCP服务器来说，连接池可以减少连接处理的时间，提高服务器的吞吐量。在Go语言中，可以使用sync.Pool来实现连接池。
// 如何实现昵 将链接放入连接池中
//4. 数据采样和分析：服务器可以采用监控工具如Prometheus和Grafana，通过对网络流量和服务状态的监控，提高服务质量和判断容量上限。
// 这个该如何实现昵 如何判断是异常连接昵，什么是异常连接昵？
//5. 开启多个端口：多个端口可以实现客户端请求的负载均衡，降低负载集中风险。
// 开多个端口为什么能降低负载集中风险昵
