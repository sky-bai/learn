package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/

func main() {

	// 开40w的链接
	for i := 0; i < 300; i++ {
		Client()
		fmt.Println("i = ", i)
	}

	OverClient()

	select {}
}

type Message struct {
	DataLen uint32 //消息的长度
	Data    []byte //消息的内容
}

func Client() {

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
	//}

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

	go func() {
		for {

			_, err = conn.Write(dataBuff.Bytes())
			if err != nil {
				fmt.Println("write error err ", err)
				return
			}

			time.Sleep(20 * time.Second)
		}
	}()

	select {}
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
	//}

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
