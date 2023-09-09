package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()

	f2, err := os.Create("mem.pprof")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 开启内存性能分析
	pprof.WriteHeapProfile(f2)

}
