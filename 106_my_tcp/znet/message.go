package znet

import "backend/bin/tcp/ziface"

var _ ziface.IMessage = (*Message)(nil)

type Message struct {
	MsgType string //消息的类型
	DataLen uint32 //消息的长度
	Data    []byte //消息的内容
}

// NewMsgPackage 创建一个Message消息包
func NewMsgPackage(msgType string, data []byte) *Message {
	return &Message{
		MsgType: msgType,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetMsgType 获取消息类型
func (msg *Message) GetMsgType() string {
	return msg.MsgType
}

// GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// SetMsgType 设计消息类型
func (msg *Message) SetMsgType(msgType string) {
	msg.MsgType = msgType
}

// SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
