package main

import (
	"fmt"
	"github.com/tarm/serial"
)

func main() {
	data()
}

func data() {

	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Print("err:\n", err)
		return
	}

	data := ""

	//不断接收buffer，直到buf='A‘,不用组装data
	for {
		// 1.每次读取一个字节
		buf := make([]byte, 1)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Print("err:\n", err)
			return
		}

		// 2.如果读取到的是'A'
		temp := string(buf[0:n])
		if temp == "A" {

			data += temp
			// 然后跳入下一次循环
			continue
		} else {
			// 拼接data
			data += temp
		}
	}

}
