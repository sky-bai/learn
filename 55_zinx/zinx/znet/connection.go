package znet

import (
	"errors"
	"fmt"
	"io"
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

	//消息管理MsgId和对应处理方法的消息管理模块
	MsgHandler ziface.IMsgHandle

	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

// NewConnection 创建连接的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	// 每个链接有它对应的回调函数，所以在创建链接的时候，将实际业务逻辑回调函数注册进来
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		MsgHandler:   msgHandler,
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

		// 创建拆包解包的对象 todo 根据协议的不同去创建不同的拆包对象
		dp := NewDataPack()

		//读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err) // 如果读取包头出错，直接退出该链接，按理说应该打印日志
			c.ExitBuffChan <- true
			continue
		}

		//拆包，得到msgid 和 datalen 放在msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			c.ExitBuffChan <- true
			continue
		}

		//根据 dataLen 读取 data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)

		//得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			msg:  msg, //将之前的buf 改成 msg
		}

		//从绑定好的消息和对应的处理方法中执行对应的Handle方法
		go c.MsgHandler.DoMsgHandler(&req)

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

// SendMsg 直接将Message数据发送数据给远程的TCP客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	// 1.先判断链接是否为空
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	// 2.拆包解包的时候，相当于都要创建个对象 将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	// 3.写回客户端
	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("Write msg id ", msgId, " error ")
		// 反正只要涉及到链接的操作，都要判断链接是否关闭，然后出错就要关闭链接
		c.ExitBuffChan <- true
		return errors.New("conn Write error")
	}

	return nil
}
