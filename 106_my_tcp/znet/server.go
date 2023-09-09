package znet

import (
	"backend/bin/tcp/ziface"
	"backend/config"
	"backend/lib/logger"
	"backend/lib/tools"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var _ ziface.IServer = (*Server)(nil)

// Server iServer 接口实现，定义一个Server服务类
type Server struct {
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	MsgHandler ziface.IMsgHandle

	//当前Server的链接管理器
	ConnMgr ziface.IConnManager

	//该Server的连接断开时的Hook函数
	onConnStop func(conn ziface.IConnection)

	//异步捕获链接关闭状态
	exitChan chan struct{}

	//心跳检测器
	hc ziface.IHeartbeatChecker
}

// 要在路由前就读取数据

// Start 开启网络服务
// 开启网络服务

func (s *Server) ListenTcpConn() {
	logger.Infof(context.Background(), "[START] Server listenner at IP: %s, Port %d, is starting", s.IP, s.Port)

	var listener *net.TCPListener

	// 1.获取一个TCP的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		logger.Errorf(context.Background(), "resolve tcp addr err: %v", err)
		return
	}

	// 2.监听服务器地址
	listener, err = net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		logger.Errorf(context.Background(), "listen %s err: %v", s.IPVersion, err)
		return
	}

	go func() {

		// 3.启动server网络连接业务
		for {
			// 3.0 如果超过最大连接数，那么就拒绝
			if uint32(s.GetConnMgr().Len()) >= config.Config.Tcp.MaxConn {
				logger.Errorf(context.Background(), "too many connections maxConn = %d", config.Config.Tcp.MaxConn)
				time.Sleep(1 * time.Second)
				continue
			}

			// 3.1 建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				logger.Warnf(context.Background(), "Accept err: %v", err)
				continue
			}

			// 3.2 增加链接
			newId := tools.RandomStr(10)

			// 3.3 处理该新连接请求的业务方法
			dealConn := NewConnection(s, conn, newId, s.MsgHandler)

			// 3.4 如果有心跳检测就启动
			if s.hc != nil {
				// 因为这里不知道设备当前状态，如果一直不发心跳消息,2分钟后就会被关闭
				s.hc.Set(dealConn, 0, 2*time.Minute)
			}

			// 3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()

	select {
	case <-s.exitChan:
		logger.Infof(context.Background(), "start listener close")
		// 下面stop方法进行关闭
		err := listener.Close()
		if err != nil {
			logger.Errorf(context.Background(), "listener close err: %v", err)
		}
		logger.Infof(context.Background(), "server stop success")
	}
}

func (s *Server) Serve() {
	s.Start()

	// Block, otherwise the listener's goroutine will exit when the main Go exits (阻塞,否则主Go退出， listenner的go将会退出)
	c := make(chan os.Signal, 1)
	// Listen for specified signals: ctrl+c or kill signal (监听指定信号 ctrl+c kill信号)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	logger.Infof(context.Background(), "server start success, pid: %d", os.Getpid())
	sig := <-c
	s.Stop()
	logger.Infof(context.Background(), "server get a sig: %v", sig)
}

func (s *Server) Start() {

	// 0.启动网络链接前server的初始化工作

	// 1.创建server退出的channel
	s.exitChan = make(chan struct{})

	// 2.启动worker工作池机制 在服务端启动之前就创建工作池
	s.MsgHandler.StartWorkerPool()

	// 3.最后启动server网络连接业务
	go s.ListenTcpConn()

}

func (s *Server) Stop() {

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
	s.exitChan <- struct{}{}
	close(s.exitChan)
}

// NewServer 创建一个服务器句柄
func NewServer() ziface.IServer {
	s := &Server{
		IPVersion:  "tcp4",
		IP:         config.Config.Tcp.Ip,
		Port:       config.Config.Tcp.Port,
		MsgHandler: NewMsgHandle(),   //msgHandler 初始化
		ConnMgr:    NewConnManager(), //创建ConnManager
		exitChan:   make(chan struct{}),
	}

	return s
}

// AddHandler Handler功能：给当前服务注册一个Handler业务方法，供客户端链接处理使用
func (s *Server) AddHandler(msgId string, handler func(request ziface.IRequest)) {
	s.MsgHandler.AddHandler(msgId, handler)

}

// GetConnMgr 得到链接管理
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

// SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.onConnStop = hookFunc
}

// StartHeartBeatWithOption StartHeartBeatWithFunc 启动心跳检测
// option 配置心跳检测
func (s *Server) StartHeartBeatWithOption(interval time.Duration, option *ziface.HeartBeatOption) {
	checker, err := NewHeartbeatChecker(interval)
	if err != nil {
		panic(err)
		return
	}

	if option != nil {
		checker.SetOnRemoteNotAlive(option.OnRemoteNotAlive)
	}

	//server绑定心跳检测器
	s.hc = checker

	// 心跳检测器绑定连接管理器
	s.hc.BindConnManger(s.ConnMgr)
}

// GetOnConnStop 得到该Server的连接断开时的Hook函数
func (s *Server) GetOnConnStop() func(ziface.IConnection) {
	return s.onConnStop
}

func (s *Server) GetMsgHandler() ziface.IMsgHandle {
	return s.MsgHandler
}

func (s *Server) GetHeartBeatChecker() ziface.IHeartbeatChecker {
	return s.hc
}
