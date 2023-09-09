package main

import (
	"backend/bin/tcp/znet"
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
	conn, err := net.Dial("tcp", "127.0.0.1:3334")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	// 点火
	SendAccOn(conn)
	//time.Sleep(1 * time.Second)
	//SendAccOn(conn)
	//// 熄火
	//SendAccOffOnline(conn)
	//time.Sleep(10 * time.Second)
	//// 点火
	//SendAccOn(conn)
	select {}
}

func SendAccOn(conn net.Conn) {
	fmt.Println("---------------------------------")
	fmt.Println("SendAccOn", time.Now().Format("2006-01-02 15:04:05"))
	msg := ""
	// 第一次发送点火在线心跳包
	msg = "#1:862582042075907:1:*,0000049C,XT,true+北京市++海淀区,V,17050F,0E2D01,016EA92C,0428B80C,03AB,1B58,000000010000,5A,4,000064,3#"

	//发封包message消息
	dp := znet.NewDataPack()
	sendData1, _ := dp.Pack(znet.NewMsgPackage(1, []byte(msg))) // 这里的ID只是个占位符，之前协议包里面没有消息ID这个字段

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
	msgHead, err := dp.Unpack(headData)
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

		fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
	}
}

func SendAccOffOnline(conn net.Conn) {
	fmt.Println("SendAccOffOnline", time.Now().Format("2006-01-02 15:04:05"))

	msg := ""
	// 第一次发送点火在线心跳包
	msg = "#1:862582042075907:1:*,0000049C,XT,true+成都市++海淀区,V,17050F,0E2D01,016EA92C,0428B80C,03AB,1B58,000000000000,5A,4,000064,3#"

	//发封包message消息
	dp := znet.NewDataPack()
	sendData1, _ := dp.Pack(znet.NewMsgPackage(1, []byte(msg))) // 这里的ID只是个占位符，之前协议包里面没有消息ID这个字段

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
	msgHead, err := dp.Unpack(headData)
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
