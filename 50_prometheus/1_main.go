package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	queueSizeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "queue_size",
		Help: "Size of the queue",
	})
)

//3. 注册指标对象
//在Prometheus服务中注册指标对象以便其可以被收集和分析：

func init() {
	prometheus.MustRegister(queueSizeGauge)
	queueSizeGauge.Set(float64(12))
}

//在init()函数中注册指标对象是标准的做法。如果您不想使用init()，也可以在其他地方注册指标对象。
//
//4. 更新指标对象
//
//使用Set()等方法来更新指标对象的值：
//
//```

//```
//
//5. 启动HTTP服务
//
//使用Prometheus HTTP处理程序公开注册的指标对象：
//
//```

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

//```
//
//6. 访问指标
//
//现在，您可以使用Prometheus进行指标查询了。在Web浏览器中打开http://localhost:8080/metrics，您将看到Prometheus在其中公开的指标。
//
//这就是Prometheus客户端库的基础。在实践中，您可以使用Prometheus客户端库来收集有关您的应用程序性能的各种指标，并使用Prometheus服务进行监视和分析。
