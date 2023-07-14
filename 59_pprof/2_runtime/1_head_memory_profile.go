package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var (
	profileName    = "memprofile.out"
	memProfileRate = 8
)

func main() {

	// 当前目录
	dir, _ := os.Getwd()

	// 获取当前目录

	filename := dir + "/59_pprof/2_runtime/" + profileName
	fmt.Println("-----", filename)
	f, err := os.Create(dir + "/59_pprof/2_runtime/" + profileName)
	if err != nil {
		fmt.Printf("memory profile creation error: %v\n", err)
		return
	}

	defer f.Close()

	time.Sleep(5 * time.Second)

	startMemProfile()
	//if err = common.Execute(op.MemProfile, 10); err != nil {
	//	fmt.Printf("execute error: %v\n", err)
	//	return
	//}
	if err := stopMemProfile(f); err != nil {
		fmt.Printf("memory profile stop error: %v\n", err)
		return
	}

}

func startMemProfile() {
	// MemProfileRate控制内存配置文件中记录和报告的内存分配比例。分析器的目标是对每个分配的MemProfileRate字节的平均分配进行采样。

	runtime.MemProfileRate = 8
}

func stopMemProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.WriteHeapProfile(f)
}
