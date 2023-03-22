package main

import (
	"mosn.io/holmes"
	"time"
)

func main() {

	// 配置规则
	h, _ := holmes.New(
		holmes.WithCollectInterval("5s"), // 指标采集时间间隔 每隔5s去采集一次
		holmes.WithDumpPath("/tmp"),      // profile保存路径 // 文件保存位置

		holmes.WithCPUDump(10, 25, 80, 2*time.Minute),                     // 配置CPU的性能监控规则
		holmes.WithMemDump(30, 25, 80, 2*time.Minute),                     // 配置Heap Memory 性能监控规则
		holmes.WithGCHeapDump(10, 20, 40, 2*time.Minute),                  // 配置基于GC周期的Heap Memory 性能监控规则
		holmes.WithGoroutineDump(500, 25, 20000, 100*1000, 2*time.Minute), //配置Goroutine数量的监控规则
	)
	// enable all
	h.EnableCPUDump().
		EnableGoroutineDump().
		EnableMemDump().
		EnableGCHeapDump().Start()
}

// 每个 Profile 都可以配置 min、diff、abs、coolDown 四个指标，含义如下:
//
// 当前指标小于 min 时，不视为异常。
//
// 当前指标大于 (100+diff)100% 历史指标，说明系统此时产生了波动，视为异常。
//
// 当前指标大于 abs (绝对值)时，视为异常。  abs 设置为最高值
//
// CPU 和 Goroutine 这两个 Profile 类型提供 Max 参数配置，基于以下考虑：
//
// CPU 的 Profiling 操作大约会有 5% 的性能损耗，所以当在 CPU 过高时，不应当进行 Profiling 操作，否则会拖垮系统。
//
// 当 Goroutine 数过大时，Goroutine Dump 操作成本很高，会进行 STW 操作，从而拖垮系统。（详情见文末参考文章）

// 对于性能监控,就要了解GC了

// 程序老是半夜崩，崩了以后就重启了，我也醒不来，现场早就丢了，不知道怎么定位

// 这压测开压之后，随机出问题，可能两小时，也可能五小时以后才出问题，这我蹲点蹲出痔疮都不一定能等到崩溃的那个时间点啊
//有些级联失败，最后留下现场并不能帮助我们准确地判断问题的根因，我们需要出问题时第一时间的现场

// Go 内置的 pprof 虽然是问题定位的神器，但是没有办法让你恰好在出问题的那个时间点，把相应的现场保存下来进行分析。特别是一些随机出现的内存泄露、CPU 抖动，等你发现有泄露的时候，可能程序已经 OOM 被 kill 掉了。而 CPU 抖动，你可以蹲了一星期都不一定蹲得到。
//
//这个问题最好的解决办法是 continuous profiling，不过这个理念需要公司的监控系统配合，在没有达到最终目标前，我们可以先向前迈一小步，看看怎么用比较低的成本来解决这个问题。
// 在公司没有完整的监控系统的时候
//从现象上，可以将出问题的症状简单分个类：
//
//cpu 抖动：有可能是模块内有一些比较冷门的逻辑，触发概率比较低，比如半夜的定时脚本，触发了以后你还在睡觉，这时候要定位就比较闹心了。
//内存使用抖动：有很多种情况会导致内存使用抖动，比如突然涌入了大量请求，导致本身创建了过多的对象。也可能是 goroutine 泄露。也可能是突然有锁冲突，也可能是突然有 IO 抖动。原因太多了，猜是没法猜出根因的。
//goroutine 数暴涨，可能是死锁，可能是数据生产完了 channel 没关闭，也可能是 IO 抖了什么的。
//CPU 使用，内存占用和 goroutine 数，都可以用数值表示，所以不管是“暴涨”还是抖动，都可以用简单的规则来表示：
//
//xx 突然比正常情况下的平均值高出了 25%
//xx 超过了模块正常情况下的最高水位线
//这两条规则可以描述大部分情况下的异常，规则一可以表示瞬时的，剧烈的抖动，之后可能迅速恢复了；规则二可以用来表示那些缓慢上升，但最终超出系统负荷的情况，例如 1s 泄露一兆内存，直至几小时后 OOM。
//
//而与均值的 diff，在没有历史数据的情况下，就只能在程序内自行收集了，比如 goroutine 的数据，我们可以每 x 秒运行一次采集，在内存中保留最近 N 个周期的 goroutine 计数，并持续与之前记录的 goroutine 数据均值进行 diff：
