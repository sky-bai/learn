package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	//发送数据
	conn.Write([]byte("第一行\n第二行\n第三行\n "))
}
