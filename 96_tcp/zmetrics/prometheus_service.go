package zmetrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
)

const (
	METRICS_ROUTE string = "/metrics"
)

var metricsServiceOnce sync.Once
var metricsInitOnce sync.Once

// PrometheusListen        string // Prometheus Metrics 服务IP和端口, 默认为 0.0.0.0:20004

func RunMetricsService(port string) (err error) {

	metricsServiceOnce.Do(func() {
		// metricsService 只启动一个服务
		go func() {
			http.Handle(METRICS_ROUTE, promhttp.Handler())
			err = http.ListenAndServe(port, nil) //多个进程不可监听同一个端口
			if err != nil {
				fmt.Println("RunMetricsService err = ", err)
				panic(err)
			}
		}()
	})

	return err
}

func InitMetrics() {

	metricsInitOnce.Do(func() {
		Metrics().connTotal = prometheus.NewGaugeVec( // 每一个指标都是gauge
			prometheus.GaugeOpts{
				Name: GaugeConnectionTotalName,
				Help: GaugeConnectionTotalHelp,
			},
			[]string{LabelAddress, LabelName}, // 该指标的标签
		)

		Metrics().taskTotal = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: GaugeTaskTotalName,
				Help: GaugeTaskTotalHelp,
			},
			[]string{LabelAddress, LabelName, LabelWorkerID},
		)

		Metrics().routerScheduleTotal = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: GaugeRouterScheduleTotalName,
				Help: GaugeRouterScheduleTotalHelp,
			},
			[]string{LabelAddress, LabelName, LabelWorkerID, LabelMsgID},
		)

		Metrics().routerScheduleDuration = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    HistogramRouterScheduleDurationName,
				Help:    HistogramRouterScheduleDurationHelp,
				Buckets: []float64{0.005, 0.01, 0.03, 0.08, 0.1, 0.5, 1.0, 5.0, 10, 100, 1000, 5000, 30000}, //单位ms,最大半分钟
			},
			[]string{LabelAddress, LabelName, LabelWorkerID, LabelMsgID},
		)

		Metrics().connDuration = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    HistogramConnDurationName,
				Help:    HistogramConnDurationHelp,
				Buckets: []float64{60000 * 30, 60000 * 30 * 2}, //单位ms,最大半分钟
			},
			[]string{LabelAddress, LabelName, LabelWorkerID, LabelMsgID},
		)

		Metrics().sum = promauto.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: "my_summary_metric",
				Help: "Summary metric example",
			},
			[]string{LabelAddress, LabelName, LabelWorkerID, LabelMsgID})

		//Register
		prometheus.MustRegister(Metrics().connTotal)
		prometheus.MustRegister(Metrics().taskTotal)
		prometheus.MustRegister(Metrics().routerScheduleTotal)
		prometheus.MustRegister(Metrics().routerScheduleDuration)
	})

}
