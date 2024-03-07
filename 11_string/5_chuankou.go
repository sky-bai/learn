package main

import (
	"fmt"
	serial "github.com/tarm/goserial"
	"goProject/zmetrics"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	//“模式，电流，电压，转速，转速百分比，温度，湿度，振动，容量使用率，母排电流”
	//模式-风机状态：2：正在运行，1：启动中，0：停机
	//电流-CT二次侧电流(A)
	//电压-电容电压(V)
	//转速-风机转速（rpm）
	//转速百分比-转速百分比
	//温度-30-40，每30s变化1°
	//湿度-50-60，每40s变化一次
	//振动-异常振动，1，0，1就是“存在”异常振动，0就是‘无’异常振动
	//容量使用率，60-70，每3s变化一次（%）
	//母排电流-母排电流（A）
	//诊断状态：默认值 健康 振动1：异常振动，电压大于160V：电容过压；温度大于60°：设备过热；
	fuzz()
	chuanKou()
	select {}

}

func chuanKou() {
	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Print("err:\n", err)
		return
	}

	buf := make([]byte, 1)
	n, err := s.Read(buf)
	if err != nil {
		fmt.Print("err:\n", err)
		return
	}
	fmt.Printf("data112222:%s\n", string(buf[0:n]))
	for {
		buf := make([]byte, 1)
		s.Read(buf)
		fmt.Printf("data112222:%s\n", string(buf[0:n]))
		if string(buf[0:n]) == "A" {
			fmt.Printf("d:%s\n", string(buf[0:n]))
			break
		}
	}
	data := ""
	//i := 1
	//不断接收buffer，直到buf='A‘,不用组装data
	for {

		buf := make([]byte, 1)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Print("err:\n", err)
			return
		}
		//fmt.Printf("data112222:%s\n", string(buf[0:n]))
		// 接收一个buf
		i := 0
		if string(buf[0:n]) == "A" { //这是data就已经是”121312“
			fmt.Printf("data:%s\n", data)
			i++
			handler(i, data)
			data = ""
			//i = 2 //这里应该对data进行处理（打印），然后将data清空
			//	continue
			//handler(data)
		} else { // if i == 2 {
			data += string(buf[0:n]) // 这里应该将每个buf串到data后面
		}

		//fmt.Printf("data11:%s\n", string(buf[0:n]))

		//time.Sleep(1 * time.Second)
	}

	// 情况3 温度大于60°
}

func fuzz() {
	zmetrics.InitMetrics()

	err := zmetrics.RunMetricsService(":9091")
	if err != nil {
		fmt.Printf("zm", err)
		return
	}

}

func handler(i int, data string) {
	dataArr := strings.Split(data, ",")
	//模式
	a0 := dataArr[0]

	// 电流
	a1 := ""
	if len(dataArr) > 0 {
		a1 = dataArr[1]
	}

	// 电压
	a2 := dataArr[2]
	// 转速
	a3 := dataArr[3]
	// 转速百分比
	a4 := dataArr[4]
	// 温度
	a5 := dataArr[5]
	// 湿度
	a6 := dataArr[6]
	// 振动
	a7 := dataArr[7]
	// 容量使用率
	a8 := dataArr[8]
	// 母排电流
	a9 := dataArr[9]

	fmt.Printf(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, "\n")
	f0, err := strconv.ParseFloat(a0, 64)
	if err != nil {
		fmt.Printf("err", err)
	}

	// mode 模式
	zmetrics.Metrics().BiliModeSetConn("local", "local", f0)

	// ElectricSetConn  电流数设置
	f1, err := strconv.ParseFloat(a1, 64)
	if err != nil {
		fmt.Printf("err", err)
	}
	zmetrics.Metrics().ElectricSetConn("local", "local", f1)

	// VdcSetConn  电压数设置
	f2, err := strconv.ParseFloat(a2, 64)
	if err != nil {
		fmt.Printf("err", err)
	}
	zmetrics.Metrics().VdcSetConn("local", "local", float64(f2))

	// RotateSetConn 转速
	f3, err := strconv.ParseFloat(a3, 64)
	if err != nil {
		fmt.Printf("err", err)
	}
	zmetrics.Metrics().RotateSetConn("local", "local", f3)

	// AimRotateSetConn 目标转速
	f4, err := strconv.ParseFloat(a4, 64)
	if err != nil {
		fmt.Printf("err", err)
	}
	zmetrics.Metrics().AimRotateSetConn("local", "local", f4)
	// TemperSetConn 温度 每30s变化1°
	f5 := 31
	if i%30 == 0 {
		rand.Seed(time.Now().UnixNano())
		f5 = rand.Intn(10) + 30
	}
	zmetrics.Metrics().TemperSetConn("local", "local", float64(f5))

	// HumiditySetConn 湿度
	f6 := 52
	if i%40 == 0 {
		rand.Seed(time.Now().UnixNano())
		f6 = rand.Intn(10) + 30
		zmetrics.Metrics().HumiditySetConn("local", "local", float64(f6))
	} else {
		zmetrics.Metrics().HumiditySetConn("local", "local", float64(100))
	}

	// VibrateSetConn 振动
	f7, err := strconv.ParseFloat(a7, 64)
	if err != nil {
		fmt.Printf("err", err)
	}
	zmetrics.Metrics().VibrateSetConn("local", "local", f7)

	// MaxPowerSetConn 容量使用率
	f8 := 65
	if i%3 == 0 {
		rand.Seed(time.Now().UnixNano())
		f8 = rand.Intn(10) + 60
	}
	zmetrics.Metrics().MaxPowerSetConn("local", "local", float64(f8))
	// MaxPowerSetConn 转速占空比
	f9, err := strconv.ParseFloat(a9, 64)
	if err != nil {
		fmt.Printf("err", err)
	}
	zmetrics.Metrics().BiliZhuanSetConn("local", "local", f9)

	// 诊断状态
	// 异常振动，电压大于150V：电容过压；温度大于60°：设备过热；
	// 情况1 振动=1
	if a7 == "1" {
		zmetrics.Metrics().BiliModeSetConn("local", "local", 1)
	}

	// 情况2 电压大于150V
	if f2 > 150 {
		zmetrics.Metrics().BiliModeSetConn("local", "local", 1)
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
