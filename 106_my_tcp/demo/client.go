package main

import (
	"backend/bin/tcp/ziface"
	"backend/bin/tcp/znet"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// 测试 设备状态改变导致心跳上传的频率改变 点火下10秒，熄火是5分钟
	// 目前就两种：000000010000 点火在线  000000000000 熄火在线  000000020000 缩时录影
	fmt.Println("Client Test ... start")

	// 3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	// 模拟设备端
	conn, err := net.Dial("tcp", "0.0.0.0:3335")
	if err != nil {
		fmt.Println("client start err, exit! err:", err)
		return
	}
	// 点火
	SendAccOn(conn)
	time.Sleep(5 * time.Second)
	SendAccOn(conn)

	//// 熄火
	//SendAccOffOnline(conn)
	//time.Sleep(10 * time.Second)
	//// 点火
	//SendAccOn(conn)
	select {}
}
func sen(conn net.Conn) {
	msg := ""
	// 第一次发送点火在线心跳包
	msg = "#1:862582042075907:1:*,0000049C,XT,true+北京市++海淀区,V,17050F,0E2D01,016EA92C,0428B80C,03AB,1B58,000000010000,5A,4,000064,3#"

	//发封包message消息
	dp := znet.NewDataPack()
	sendData1, _ := dp.Pack(znet.NewMsgPackage("1", []byte(msg))) // 这里的ID只是个占位符，之前协议包里面没有消息ID这个字段

	_, err := conn.Write(sendData1)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}
}

func SendAccOn(conn net.Conn) {
	fmt.Println("---------------------------------")
	fmt.Println("SendAccOn", time.Now().Format("2006-01-02 15:04:05"))
	msg := ""
	// 第一次发送点火在线心跳包
	msg = "#1:862582042075907:1:*,0000049C,XT,true+北京市++海淀区,V,17060F,0E3330,016EA92C,0428B80C,03AB,1B58,000000010000,5A,4,000064,3#"

	//发封包message消息
	dp := znet.NewDataPack()
	sendData1, _ := dp.Pack(znet.NewMsgPackage("XT", []byte(msg))) // 这里的ID只是个占位符，之前协议包里面没有消息ID这个字段

	_, err := conn.Write(sendData1)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}

	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	if err != nil {
		fmt.Println("read head error", err)
		return
	}

	//将headData字节流 拆包到msg中
	msgHead, err := Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return
	}

	if msgHead.GetDataLen() > 0 {
		//msg 是有data数据的，需要再次读取data数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		//根据dataLen从io中读取字节流
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("server unpack data err:", err)
			return
		}

		fmt.Println("==> Recv Msg: ID=", msg.MsgType, ", len=", msg.DataLen, ", data=", string(msg.Data))
	}
}

func Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head的信息，得到dataLen和msgID
	msg := &znet.Message{}

	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}

func SendAccOffOnline(conn net.Conn) {
	fmt.Println("SendAccOffOnline", time.Now().Format("2006-01-02 15:04:05"))

	msg := ""
	// 第一次发送点火在线心跳包
	msg = "#1:862582042075907:1:*,0000049C,XT,true+成都市++海淀区,V,17050F,0E2D01,016EA92C,0428B80C,03AB,1B58,000000000000,5A,4,000064,3#"

	//发封包message消息
	dp := znet.NewDataPack()
	sendData1, _ := dp.Pack(znet.NewMsgPackage("1", []byte(msg))) // 这里的ID只是个占位符，之前协议包里面没有消息ID这个字段

	_, err := conn.Write(sendData1)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}

	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	if err != nil {
		fmt.Println("read head error")
		return
	}

	//将headData字节流 拆包到msg中
	msgHead, err := Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return
	}

	if msgHead.GetDataLen() > 0 {
		//msg 是有data数据的，需要再次读取data数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		//根据dataLen从io中读取字节流
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("server unpack data err:", err)
			return
		}

		//fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
	}
}
