package ziface

// 现在我们就给用户提供一个自定义的conn处理业务的接口吧，很显然，
// 我们不能把业务处理的方法绑死在type HandFunc func(*net.TCPConn, []byte, int) error这种格式中，
// 我们需要定一些interface{}来让用户填写任意格式的连接处理业务方法。

// 我们现在需要把客户端请求的连接信息 和 请求的数据，放在一个叫Request的请求类里，这样的好处是我们可以从Request里得到全部客户端的请求信息，
// 也为我们之后拓展框架有一定的作用，一旦客户端有额外的含义的数据信息，都可以放在这个Request里。可以理解为每次客户端的全部请求数据，Zinx都会把它们一起放到一个Request结构体里。

/*
IRequest 接口：
实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
	GetMsgID() uint32           //获取请求的消息ID
}

// 不难看出，当前的抽象层只提供了两个Getter方法，所以有个成员应该是必须的，一个是客户端连接，一个是客户端传递进来的数据，当然随着Zinx框架的功能丰富，这里面还应该继续添加新的成员。
