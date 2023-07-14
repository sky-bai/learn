package ziface

// IMessage 将请求的一个消息封装到message中，定义抽象层接口
type IMessage interface {
	SetMsgType(string)  //设计消息类型
	GetMsgType() string //获取消息类型

	SetData([]byte)     //设计消息内容
	GetData() []byte    //获取消息内容
	GetDataLen() uint32 //获取消息数据段长度
}
