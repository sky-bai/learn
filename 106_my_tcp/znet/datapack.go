package znet

import (
	"backend/bin/tcp/ziface"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var _ ziface.IDataPack = (*DataPack)(nil)

var MsgNil = errors.New("msg data is nil")

// DataPack 封包拆包类实例，暂时不需要成员
type DataPack struct{}

// NewDataPack 封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	return 4
}

// Pack 封包方法(压缩数据)
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 拆包方法(解压数据) 拆包的时候只需要把head信息读出来，然后再根据head信息里的data长度，再进行一次读
func (dp *DataPack) Unpack(c ziface.IConnection) (ziface.IMessage, error) {

	// 1.读取客户端的4个字节
	headData := make([]byte, dp.GetHeadLen()) // 这里只是将4个字节读取出来
	_, err := io.ReadFull(c.GetTCPConnection(), headData)
	if err != nil {
		return nil, err
	}

	// 2.创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(headData)

	msg := &Message{}

	// 3.读取dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 4.判断是否有数据
	if msg.GetDataLen() <= 0 {
		return nil, MsgNil
	}

	// 5.根据dataLen读取data，放在msg.Data中
	var data []byte
	data = make([]byte, msg.GetDataLen())
	if _, err = io.ReadFull(c.GetTCPConnection(), data); err != nil {
		return nil, err
	}

	// 5.设置拆包对象的data
	msg.SetData(data)

	return msg, nil
}
