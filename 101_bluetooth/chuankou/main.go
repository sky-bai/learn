package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"log"
	"time"
)

func main() {

	// 1.设置串口设备名称 和 波特率
	c := &serial.Config{Name: "设备名", Baud: 115200}

	// 2.打开串口
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	// 3.延迟1秒接收信号
	time.Sleep(1 * time.Second)

	// 4.然后一直读取信号
	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal("s read err:\n", err)
		}

		fmt.Printf("Read %d Bytes\r\n", n)
		for i := 0; i < n; i++ {
			fmt.Printf("buf[%d]=%c\r\n", i, buf[i])
		}
	}

}

// 我这边的问题：
// 1.我没有串口的设备，刚刚用公司的usb设备发现都不能串口通信。
// 2.我没有windows环境，卧槽，用了三个库，都只支持linux和mac，windows都不支持。

// 先看看这段程序能不能在你电脑上接收到设备的数据
// 在Windows电脑上，可以使用以下命令来列出所有已识别的串口设备：
// 1.wmic path Win32_SerialPort get DeviceID, Caption
// 2.mode
// 在Windows系统上，串口设备名称通常以COM开头，后面跟着一个数字，例如COM1、COM2等。

// 然后把设备名填入上方的 设备名 中，看看能不能读取到数据。

// ---------------------------------------------------------
//
// 问题1.我不知道读出来的数据有什么。你需要展示什么数据。
// 问题2.数据格式是什么，第一位是表示什么数据,第二位又表示什么。
// 然后需要你这边帮忙一下，看看获取的数据有什么，然后格式是什么。
