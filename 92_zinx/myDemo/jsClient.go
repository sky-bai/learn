package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"learn/55_zinx/zinx/znet"
	"net"
	"time"
)

/*
模拟客户端
*/

// lim := syscall.Rlimit{
//		Cur: math.MaxInt64,
//		Max: math.MaxInt64,
//	}
//	syscall.Setrlimit(syscall.RLIMIT_NOFILE, &lim)

// // go语言客户端 net.Dial 报错 socket: too many open files ,这是因为什么昵

func main() {

	Client()

	//OverClient()

	select {}
}

type Message struct {
	DataLen uint32 //消息的长度
	Data    []byte //消息的内容
}

func Client() {
	//
	conn, err := net.Dial("tcp", "127.0.0.1:3334")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}
	dataBuff := bytes.NewBuffer([]byte{})
	//var da Message
	//
	//da.Data = []byte("yJEqGqkbHZ<<<#1:862582042075907:1:*,1,GPS,1682329329182,wgs84,40047861:116289569:1682329325182:-782:46:2299:9600:1500:27:0|40047858:116289577:1682329326178:-778:46:2338:9600:1500:27:0|40047855:116289580:1682329327178:-778:46:2414:9600:1500:27:0|40047849:116289576:1682329328178:-778:40:2327:10300:1500:27:0|40047844:116289582:1682329329182:-782:40:2328:10300:1500:27:0,1#")
	//da.DataLen = uint32(len(da.Data))
	////写dataLen
	//if err := binary.Write(dataBuff, binary.LittleEndian, da.DataLen); err != nil {
	//	fmt.Println("write dataLen error err ", err)
	//}
	//
	////写data数据
	//if err := binary.Write(dataBuff, binary.LittleEndian, da.Data); err != nil {
	//	fmt.Println("write dataLen error err ", err)
	//
	//}

	var dataXt Message

	dataXt.Data = []byte("NGLGILxwBb<<<#1:869497050200318:1:*,0000027D,XT,true+++,V,170504,100819,00000000,00000000,0000,0000,000000010000,5A,4,000064,100#")
	dataXt.Data = []byte("FAhaQQpwfL<<<#1:803278210401627:1:*,1,GPS,1685519752617,wgs84,31089138:121505901:1685519748618:4382:0:4357:5700:496:34:0|31089138:121505901:1685519749616:4384:0:4357:5700:495:34:0|31089138:121505901:1685519750616:4384:28:4357:5700:495:34:0|31089138:121505901:1685519751617:4383:0:4357:5700:496:34:0|31089138:121505901:1685519752617:4383:32:4357:5700:496:34:0,1#")
	dataXt.DataLen = uint32(len(dataXt.Data))

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, dataXt.DataLen); err != nil {
		fmt.Println("write dataLen error err ", err)
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, dataXt.Data); err != nil {
		fmt.Println("write dataLen error err ", err)

	}

	for {
		dp := znet.NewDataPack()
		_, err = conn.Write(dataBuff.Bytes())
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
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

		time.Sleep(5 * time.Second)
	}
}

func OverClient() {

	conn, err := net.Dial("tcp", "127.0.0.1:7777") // go 语言 net.Dial 报错 socket: too many open files ,这是因为什么昵
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}
	dataBuff := bytes.NewBuffer([]byte{})
	//var da Message
	//
	//da.Data = []byte("yJEqGqkbHZ<<<#1:862582042075907:1:*,1,GPS,1682329329182,wgs84,40047861:116289569:1682329325182:-782:46:2299:9600:1500:27:0|40047858:116289577:1682329326178:-778:46:2338:9600:1500:27:0|40047855:116289580:1682329327178:-778:46:2414:9600:1500:27:0|40047849:116289576:1682329328178:-778:40:2327:10300:1500:27:0|40047844:116289582:1682329329182:-782:40:2328:10300:1500:27:0,1#")
	//da.DataLen = uint32(len(da.Data))
	////写dataLen
	//if err := binary.Write(dataBuff, binary.LittleEndian, da.DataLen); err != nil {
	//	fmt.Println("write dataLen error err ", err)
	//}
	//
	////写data数据
	//if err := binary.Write(dataBuff, binary.LittleEndian, da.Data); err != nil {
	//	fmt.Println("write dataLen error err ", err)
	//
	//} // 之前第一家有 用的是jb家的suona

	var dataXt Message

	dataXt.Data = []byte("NGLGILxwBb<<<#1:869497050200318:1:*,0000027D,XT,true+++,V,170504,100819,00000000,00000000,0000,0000,000000010000,5A,4,000064,100#")
	dataXt.DataLen = uint32(len(dataXt.Data))

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, dataXt.DataLen); err != nil {
		fmt.Println("write dataLen error err ", err)
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, dataXt.Data); err != nil {
		fmt.Println("write dataLen error err ", err)

	}

	_, err = conn.Write(dataBuff.Bytes())
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}
	return
}
