package main

import (
	"fmt"
	"mosn.io/holmes"
	"mosn.io/holmes/reporters/pyroscope_reporter"
	"os"
	"time"
)

func main() {

	UpstreamAddress := ""
	ApplicationName := ""

	cfg := pyroscope_reporter.RemoteConfig{
		// 上报地址
		UpstreamAddress: UpstreamAddress,
		// 上报请求超时
		UpstreamRequestTimeout: 3 * time.Second,
	}

	host, _ := os.Hostname()
	// 要填自己pod的hostname 这样作为tag好排查
	tags := map[string]string{"host": host}
	pReporter, err := pyroscope_reporter.NewPyroscopeReporter(ApplicationName, tags, cfg, holmes.NewStdLogger())
	if err != nil {
		fmt.Printf("NewPyroscopeReporter error %v\n", err)
		return
	}

	h, err := holmes.New(
		holmes.WithProfileReporter(pReporter),
		holmes.WithDumpPath("/tmp"),
		holmes.WithMemoryLimit(100*1024*1024), // 100MB
		holmes.WithCPUCore(2),                 // 双核
		holmes.WithCPUDump(20, 100, 150, time.Minute*2),
		holmes.WithMemDump(50, 100, 800, time.Minute),
		holmes.WithGoroutineDump(200, 100, 5000, 200*5000, 1*time.Minute),
		holmes.WithCollectInterval("1s"),
	)
	h.EnableCPUDump().
		EnableGoroutineDump().
		EnableMemDump().
		Start()
	defer h.Stop()
	// your code goes here
}
