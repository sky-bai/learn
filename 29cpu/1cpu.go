package main

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	//cpuInfos, err := cpu.Info()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, v := range cpuInfos {
	//	fmt.Println("---", v)
	//	// 负载
	avg, err := load.Avg()
	if err != nil {
		fmt.Println("load avg err:", err)
		return
	}
	fmt.Println("load avg:", avg)
	//}
	//for {
	//	percent, err := cpu.Percent(time.Second, false)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println("cpu percent:", percent)
	//}
	// 内存
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("memInfo:", memInfo)
}
