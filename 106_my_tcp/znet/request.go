package znet

import "backend/bin/tcp/ziface"

var _ ziface.IRequest = (*Request)(nil)

// 我们现在需要把客户端请求的连接信息 和 请求的数据，放在一个叫Request的请求类里，这样的好处是我们可以从Request里得到全部客户端的请求信息，
// 一旦客户端有额外的含义的数据信息，都可以放在这个Request里。可以理解为每次客户端的全部请求数据，都会把它们一起放到一个Request结构体里。

type Request struct {
	conn ziface.IConnection //已经和客户端建立好的 链接
	msg  ziface.IMessage    //客户端请求的数据
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgType 获取请求的消息类型
func (r *Request) GetMsgType() string {
	return r.msg.GetMsgType()
}

// SetMsgType 设置请求的消息类型
func (r *Request) SetMsgType(msgType string) {
	r.msg.SetMsgType(msgType)
}
