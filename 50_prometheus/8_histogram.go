package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	dto "github.com/prometheus/client_model/go"
	"math"

	"github.com/prometheus/client_golang/prometheus"
)

// A Histogram counts individual observations from an event or sample stream in configurable static buckets (or in dynamic sparse buckets as part of the experimental Native Histograms, see below for more details). Similar to a Summary, it also provides a sum of observations and an observation count.
//
// On the Prometheus server, quantiles can be calculated from a Histogram using the histogram_quantile PromQL function.
//
// Note that Histograms, in contrast to Summaries, can be aggregated in PromQL (see the documentation for detailed procedures). However, Histograms require the user to pre-define suitable buckets, and they are in general less accurate. (Both problems are addressed by the experimental Native Histograms. To use them, configure a NativeHistogramBucketFactor in the HistogramOpts. They also require a Prometheus server v2.40+ with the corresponding feature flag enabled.)
//
// The Observe method of a Histogram has a very low performance overhead in comparison with the Observe method of a Summary.
//
// To create Histogram instances, use NewHistogram.
func main() {
	temps := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "pond_temperature_celsius",
		Help:    "The temperature of the frog pond.", // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(20, 5, 5),  // 5 buckets, each 5 centigrade wide.
	})

	//go func() {
	// Simulate some observations.
	for i := 0; i < 1000; i++ {
		temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	}
	//}()

	// Just for demonstration, let's check the state of the histogram by
	// (ab)using its Write method (which is usually only used by Prometheus
	// internally).
	metric := &dto.Metric{}
	temps.Write(metric)
	fmt.Println(proto.MarshalTextString(metric))

}
