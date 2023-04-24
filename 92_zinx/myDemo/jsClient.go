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

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会

	conn, err := net.Dial("tcp", "127.0.0.1:3334")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	dataBuff := bytes.NewBuffer([]byte{})
	var da Message

	da.Data = []byte("yJEqGqkbHZ<<<#1:862582042075907:1:*,1,GPS,1682329329182,wgs84,40047861:116289569:1682329325182:-782:46:2299:9600:1500:27:0|40047858:116289577:1682329326178:-778:46:2338:9600:1500:27:0|40047855:116289580:1682329327178:-778:46:2414:9600:1500:27:0|40047849:116289576:1682329328178:-778:40:2327:10300:1500:27:0|40047844:116289582:1682329329182:-782:40:2328:10300:1500:27:0,1#")
	da.DataLen = uint32(len(da.Data))
	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, da.DataLen); err != nil {
		fmt.Println("write dataLen error err ", err)
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, da.Data); err != nil {
		fmt.Println("write dataLen error err ", err)

	}
	for {

		_, err := conn.Write(dataBuff.Bytes())
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}

type Message struct {
	DataLen uint32 //消息的长度
	Data    []byte //消息的内容
}
