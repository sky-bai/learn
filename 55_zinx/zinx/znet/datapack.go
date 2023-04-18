package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"learn/55_zinx/zinx/utils"
	"learn/55_zinx/zinx/ziface"
)

// DataPack 封包拆包类实例，暂时不需要成员
type DataPack struct{}

// NewDataPack 封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//Id uint32(4字节) +  DataLen uint32(4字节)
	return 8
}

// Pack 封包方法(压缩数据)
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}
	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 在这段代码中，结构体 Message 的字段顺序是 Id、DataLen、Data，在解析数据包时，根据顺序读取 DataLen、Id 字段的数据，解析完 DataLen 和 Id 之后， Data 字段的内容还没有读取。
	//
	// 在这里，binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen) 读取的数据长度只是 DataLen 字段所占用的长度，读取时会将数据解析为 uint32 类型，随后存入 msg.DataLen 字段中。
	// 相当于是我存入了一个 uint32 类型的数据，读取出的也是一个 uint32 类型的数据，这个数据就是 DataLen 字段的值。

	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断dataLen的长度是否超出我们允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("Too large msg data recieved")
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}

// 需要注意的是整理的Unpack方法，因为我们从上图可以知道，我们进行拆包的时候是分两次过程的，
// 第二次是依赖第一次的dataLen结果，所以Unpack只能解压出包头head的内容，得到msgId 和 dataLen。
// 之后调用者再根据dataLen继续从io流中读取body中的数据。
