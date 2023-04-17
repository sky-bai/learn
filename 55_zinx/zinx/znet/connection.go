package znet

import (
	"fmt"
	"learn/55_zinx/zinx/ziface"
	"net"
)

// 链接对象

type Connection struct { // 对象的属性

	//当前连接的原生socket TCP套接字
	Conn *net.TCPConn
	//当前连接的ID 也可以称作为SessionID，ID全局唯一 todo 这个如何赋值
	ConnID uint32
	//当前连接的关闭状态
	isClosed bool

	//该连接的处理方法router
	Router ziface.IRouter

	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

// NewConnection 创建连接的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	// 每个链接有它对应的回调函数，所以在创建链接的时候，将实际业务逻辑回调函数注册进来
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1), // 这里为什么不用make(chan struct{}, 1)呢？ 因为struct{}{}占用的内存空间为0，而bool占用的内存空间为1 所以为什么不用struct{}{}呢？ 因为struct{}{}不能赋值，而bool可以赋值为true或者false 所以为什么不用bool呢？ 因为bool占用的内存空间为1，而struct{}{}占用的内存空间为0
	}

	return c
}

// StartReader  处理conn读数据的Goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is  running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 1.读取我们最大的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			// 读完了就退出 针对数据读完的情况
			c.ExitBuffChan <- true
			continue
		}

		// 2.得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			data: buf,
		}

		// 3.从路由Routers 中找到注册绑定Conn的对应Handle
		go func(request ziface.IRequest) { // 把请求交给路由对象 让他去处理
			//执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		// 这里我们在conn读取完客户端数据之后，将数据和conn封装到一个Request中，作为Router的输入数据。
		//
		// 然后我们开启一个goroutine去调用给Zinx框架注册好的路由业务。
	}
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {

	// 1.开启处理该链接读取到客户端数据之后然后请求业务
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			//得到退出消息，不再阻塞
			// 这里监听到消息就return,可是并没有关闭连接，这里是不是有问题？ 有问题，这里应该调用Stop方法
			return
		}
	}
}

// Stop 停止连接，结束当前连接状态M
func (c *Connection) Stop() {
	//1. 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//TODO Connection Stop() 如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用

	// 关闭所有资源
	// 关闭socket链接
	c.Conn.Close()

	//通知从缓冲队列读数据的业务，该链接已经关闭
	c.ExitBuffChan <- true

	//关闭该链接全部管道
	close(c.ExitBuffChan)
}

// GetTCPConnection 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
