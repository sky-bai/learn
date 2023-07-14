package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/currantlabs/ble"
	"github.com/currantlabs/ble/examples/lib/dev"
)

func main() {
	f2()
	select {}
}

func f2() {

	// 打开 BLE 设备
	d, err := dev.NewDevice("B.O.W")
	if err != nil {
		log.Fatalf("Failed to open BLE device: %s", err)
	}
	defer d.Stop()

	// 扫描设备
	ble.SetDefaultDevice(d)
	ble.Scan(context.Background(), false, advHandler, nil)

	// 等待一段时间，以便获取设备信号强度
	time.Sleep(5 * time.Second)

}

// 广告处理程序
func advHandler(a ble.Advertisement) {
	fmt.Printf("Peripheral Discovered: %s\n", a.Address())
	fmt.Printf("  Name: %s\n", a.LocalName())
	fmt.Printf("  RSSI: %d dBm\n", a.RSSI())
}

// gatt.
//	device, err := gatt.NewDevice(nil)
//	if err != nil {
//		log.Fatalf("Failed to open device: %s", err)
//	}
//	// 您好，现在我有一个硬件设备，然后和我自己的电脑连的同一个热点，现在我想在我的windows系统电脑上使用go语言编写获取设备蓝牙信号的程序，你能教教我吗？先说声谢谢了
//	device.Handle(gatt.PeripheralDiscovered(func(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
//		fmt.Printf("Peripheral Discovered: %s\n", p.ID())
//		fmt.Printf("  Name: %s\n", p.Name())
//		fmt.Printf("  RSSI: %d\n", rssi)
//	}))
//
//	device.Init(func(d gatt.Device, s gatt.State) {
//		if s == gatt.StatePoweredOn {
//			fmt.Println("Scanning for devices...")
//			d.Scan([]gatt.UUID{}, false)
//			time.Sleep(5 * time.Second)
//			d.StopScanning()
//		} else {
//			log.Fatalf("Failed to initialize device, state: %s", s)
//		}
//	})
