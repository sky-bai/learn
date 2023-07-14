package main

//
//import (
//	"fmt"
//	"github.com/microsoft/go-winio/pkg/devicereg"
//	"github.com/microsoft/go-winio/pkg/guid"
//	"log"
//)
//
//func main() {
//	// 蓝牙设备的 GUID
//	guidBluetoothDevice, err := guid.FromString("{EA3B5B82-26EE-450E-B0D8-D26FE9A61F00}")
//	if err != nil {
//		log.Fatalf("Error creating Bluetooth GUID: %v", err)
//	}
//
//	// 查询蓝牙设备
//	devices, err := devicereg.QueryDeviceByGUID(guidBluetoothDevice)
//	if err != nil {
//		log.Fatalf("Error querying Bluetooth devices: %v", err)
//	}
//
//	// 遍历搜索到的蓝牙设备并打印信息
//	for _, device := range devices {
//		deviceName, err := device.DeviceName()
//		if err != nil {
//			log.Printf("Error getting device name: %v", err)
//			continue
//		}
//
//		fmt.Printf("Device Name: %s\n", deviceName)
//
//		// 获取蓝牙信号强度
//		rssi, err := getBluetoothSignalStrength(device)
//		if err != nil {
//			log.Printf("Error getting Bluetooth signal strength: %v", err)
//		} else {
//			fmt.Printf("Bluetooth Signal Strength (RSSI): %d dBm\n", rssi)
//		}
//
//		fmt.Println()
//	}
//}
//
//// getBluetoothSignalStrength 获取蓝牙设备的信号强度 (RSSI)
//func getBluetoothSignalStrength(device *devicereg.DeviceInfo) (int, error) {
//	// 连接设备
//	err := device.Connect()
//	if err != nil {
//		return 0, fmt.Errorf("Error connecting to device: %w", err)
//	}
//	defer device.Disconnect()
//
//	// 获取信号强度
//	rssi, err := device.SignalStrength()
//	if err != nil {
//		return 0, fmt.Errorf("Error getting signal strength: %w", err)
//	}
//
//	return rssi, nil
//}
