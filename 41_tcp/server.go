package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

// 出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此需要一种机制来告诉接收方数据包的大小，这种机制就是”包头”。
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	// 去掉包头，获取消息内容 // 如果数据包超过了缓存区的大小了昵，那么就会出现粘包的情况
	return string(pack[4:]), nil
}

// 如果是tcp服务器的话，什么时候需要解决粘包和毡包问题呢？
// 1. 服务器端接收到的数据包大小不确定
// 2. 服务器端接收到的数据包大小不确定，且数据包的大小超过了缓冲区的大小

// 为什么用tcp不用http昵？
// 1. http是基于tcp的，http的请求和响应都是基于tcp的

// 为什么用tcp不用udp昵？
// 1. udp是无连接的，不需要建立连接，但是tcp是面向连接的，需要建立连接
// 2. udp是不可靠的，不保证数据的可靠传输，但是tcp是可靠的，保证数据的可靠传输
// 3. udp是面向报文的，不需要拼接数据包，但是tcp是面向字节流的，需要拼接数据包
// 4. udp是没有拥塞控制的，但是tcp有拥塞控制

// 为什么用tcp不用websocket昵？
// 1. websocket是基于tcp的，websocket的请求和响应都是基于tcp的
