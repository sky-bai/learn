package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":6060", http.DefaultServeMux)
}
