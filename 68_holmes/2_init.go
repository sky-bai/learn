package main

import (
	mlog "mosn.io/pkg/log"
	"time"

	"mosn.io/holmes"
)

func main() {

	h := initHolmes()

	// start the metrics collect and dump loop
	h.Start()

}
func initHolmes() *holmes.Holmes {
	h, _ := holmes.New(

		holmes.WithCollectInterval("15s"),                                    // 1.间隔几秒收集一次数据 线上一般设置成15s
		holmes.WithDumpPath("./tmp"),                                         // 2.数据存放的目录
		holmes.WithCPUDump(10, 25, 80, time.Minute),                          // 3.设置cpu dump的参数
		holmes.WithCPUMax(85),                                                // 4.WithCPUMax 当cpu使用率大于Max, holmes会跳过dump操作，以防拖垮系统。
		holmes.WithGoroutineDump(10, 25, 2000, 10*1000, time.Minute),         // 5.设置goroutine dump的参数
		holmes.WithMemDump(30, 25, 80, time.Minute),                          // 6.设置内存dump的参数
		holmes.WithGCHeapDump(10, 20, 40, time.Minute),                       // 7.设置GCHeap dump的参数
		holmes.WithTextDump(),                                                // 8.以文本格式保存profile内容。
		holmes.WithLogger(holmes.NewFileLog("./tmp/holmes.log", mlog.DEBUG)), // 9.保存holmes 日志
	)
	h.EnableCPUDump()       // 开启cpu dump
	h.EnableGoroutineDump() // 开启goroutine dump
	h.EnableMemDump()       // 开启内存dump
	h.EnableGCHeapDump()    // 开启gc dump
	return h
}

// WithCPUDump(10, 25, 80, time.Minute) 会在满足以下条件时dump profile cpu usage > 10% && cpu usage > 125% * previous cpu usage recorded or cpu usage > 80%.
// time.Minute 是两次dump操作之间最小时间间隔，避免频繁profiling对性能产生的影响。

// WithGoroutineDump(10, 25, 2000, 100*1000, time.Minute) 当goroutine指标满足以下条件时，将会触发dump操作。 current_goroutine_num > 10 && current_goroutine_num < 100*1000 && current_goroutine_num > 125% * previous_average_goroutine_num or current_goroutine_num > 2000.
// time.Minute 是两次dump操作之间最小时间间隔，避免频繁profiling对性能产生的影响。
// 当应用所启动的goroutine number大于Max 时，holmes会跳过dump操作，因为当goroutine number很大时， dump goroutine profile操作成本很高（STW && dump），有可能拖垮应用。当Max=0 时代表没有限制。

// WithMemDump(10, 25, 80, time.Minute) 会在满足以下条件时抓取heap profile memory usage > 10% && memory usage > 125% * previous memory usage or memory usage > 80%，
// time.Minute 是两次dump操作之间最小时间间隔，避免频繁profiling对性能产生的影响。
