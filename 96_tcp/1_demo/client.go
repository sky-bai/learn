package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	ip          = flag.String("ip", "127.0.0.1", "server IP")
	connections = flag.Int("conn", 1, "number of tcp connections")
)

// 参数 默认值 说明
// 每一个链接都是一个协程

func main() {
	flag.Parse()
	fmt.Println("ip:", ip)
	fmt.Println("connections:", connections)
	addr := *ip + ":8972"
	log.Printf("连接到 %s", addr)
	var conns []net.Conn
	for i := 0; i < *connections; i++ {
		c, err := net.DialTimeout("tcp", addr, 10*time.Second)
		if err != nil {
			fmt.Println("failed to connect", i, err)
			i--
			continue
		}
		conns = append(conns, c)
		time.Sleep(time.Millisecond)
	}
	defer func() {
		for _, c := range conns {
			c.Close()
		}
	}()
	log.Printf("完成初始化 %d 连接", len(conns))
	tts := time.Second
	if *connections > 100 {
		tts = time.Millisecond * 5
	}
	for {
		for i := 0; i < len(conns); i++ {
			time.Sleep(tts)
			conn := conns[i]
			conn.Write([]byte("hello world\r\n"))
		}
	}
}
