package ziface

// V0.1版本我们已经实现了一个基础的Server框架，现在我们需要对客户端链接和不同的客户端链接所处理的不同业务再做一层接口封装，当然我们先是把架构搭建起来。
// 现在在 ziface 下创建一个属于链接的接⼝文件 iconnection.go ，当然他的实现⽂文件我们放在 znet 下的 connection.go 中

import "net"

// IConnection 定义连接⼝
type IConnection interface {
	// Start 启动连接，让当前连接开始工作
	Start()
	// Stop 停⽌止连接，结束当前连接状态M
	Stop()
	// GetConnID 从当前连接获取原始的socket TCPConn GetTCPConnection() *net.TCPConn //获取当前连接ID
	GetConnID() uint32 //获取远程客户端地址信息 RemoteAddr() net.Addr

	GetTCPConnection() *net.TCPConn

	// SendMsg 直接将Message数据发送数据给远程的TCP客户端
	SendMsg(msgId uint32, data []byte) error
}

// HandFunc 定义一个统一处理链接业务的接口 包含链接 请求端的数据 以及数据的长度
type HandFunc func(*net.TCPConn, []byte, int) error

// 该接⼝的一些基础方法，代码注释已经介绍的很清楚，这⾥先简单说明⼀一个HandFunc这个函数类型， 这个是所有conn链接在处理业务的函数接口，
// 第一参数是socket原⽣链接，第⼆个参数是客户端请求的数据，第三个参数是客户端请求的数据长度。这样，如果我们想要指定一个conn的处理业务，只要定义一个HandFunc类型的函数，然后和该链接绑定就可以了。

// 现在我们已经将拆包的功能集成到Zinx中了，但是使用Zinx的时候，如果我们希望给用户返回一个TLV格式的数据，总不能每次都经过这么繁琐的过程，所以我们应该给Zinx提供一个封包的接口，供Zinx发包使用。
