package zmetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

var _metrics *zMetrics
var _metricsOnce sync.Once

type zMetrics struct {
	// 链接总数
	connTotal *prometheus.GaugeVec //[address, name]:ConnTotal
	// 每个workId处理的任务总数
	taskTotal *prometheus.GaugeVec //[address, name, workerID]:TaskTotal

	// 每一个链接处理的消息总数
	messageTotal *prometheus.GaugeVec //[address, name, connId]:MessageTotal

	// 当前请求的qps
	qpsTotal *prometheus.GaugeVec //[address, name, connId]:MessageTotal

	// 路由Router调度的Handler次数
	routerScheduleTotal *prometheus.GaugeVec //[address, name, workerID, MsgID]:RouterScheduleTotal
	// 路由Router调度的Handler耗时
	routerScheduleDuration *prometheus.HistogramVec //[address, name, workerID, MsgID]:RouterScheduleDuration

	// 连接时长
	connDuration *prometheus.HistogramVec

	sum *prometheus.SummaryVec
}

// 每一个链接处理的消息总数

// Metrics 获取单例
func Metrics() *zMetrics {
	_metricsOnce.Do(func() {
		_metrics = new(zMetrics)
	})
	return _metrics
}

// IncConn 链接数量累加
func (m *zMetrics) IncConn(serverAddress, serverName string) {
	m.connTotal.WithLabelValues(serverAddress, serverName).Inc()

}

// DecConn 链接数量累减
func (m *zMetrics) DecConn(serverAddress, serverName string) {
	m.connTotal.WithLabelValues(serverAddress, serverName).Dec()

}

// IncTask 任务数量累加
func (m *zMetrics) IncTask(address, name, workerID string) {
	m.taskTotal.WithLabelValues(address, name, workerID).Inc()
}

func (m *zMetrics) IncRouterSchedule(address, name, workerID, msgID string) {
	m.routerScheduleTotal.WithLabelValues(address, name, workerID, msgID).Inc()
}

func (m *zMetrics) ObserveRouterScheduleDuration(address, name, workerID, msgID string, duration time.Duration) {
	m.routerScheduleDuration.With(
		prometheus.Labels{
			LabelAddress:  address,
			LabelName:     name,
			LabelWorkerID: workerID,
			LabelMsgID:    msgID,
		}).Observe(duration.Seconds() * 1000)
}

func (m *zMetrics) SumConn(serverAddress, serverName string) {
	m.sum.WithLabelValues(serverAddress, serverName).Observe(1)

}
