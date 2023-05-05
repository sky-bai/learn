package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	opsQueued := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "ops_queued",
		Help:      "Number of blob storage operations waiting to be processed.",
	})

	// our_company_blob_storage_ops_queued 输入完整的指标名称

	// 在Prometheus中，Namespace用于对系统中的监控指标进行分类和命名，通常用于区分部署相同软件的不同环境或子系统。
	// 它允许同一种指标的不同环境或子系统具有相同的名称，这些指标能够被Prometheus以不同的方式区分和聚合。
	// 在你的例子中，Namespace是指“our_company”，代表所属公司或组织。

	// prometheus中gauge类型指标对象的介绍
	prometheus.MustRegister(opsQueued)
	go func() {
		for {

			opsQueued.Add(1)

			time.Sleep(2 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

//

// 在Prometheus中，Gauge是一种指标类型，表示某些可变数量的值。
// 它通常用于测量诸如温度、CPU占用率、磁盘使用、内存使用、请求处理时间等变量，这些变量的值可能会上升或下降，但它们的当前值并不直接反映趋势或速率。
//
// 一个Gauge类型的指标对象由一个浮点数值组成，它表示所测量的数量的当前值。Gauge的值可以递增或递减，并且可以在任何时刻进行设置。
//
// Gauge类型指标对象还提供了一些方法来对其值进行增加、减少和设置操作，例如Add()、Inc()、Dec()和Set()。
// 这些方法允许应用程序通过直接修改计数器的值来更新Gauge指标对象的当前值。
