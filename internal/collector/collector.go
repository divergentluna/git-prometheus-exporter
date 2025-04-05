package collector

import "github.com/prometheus/client_golang/prometheus"

type GitMetricsCollector interface {
	Describe(ch chan<- *prometheus.Desc)
	Collect(ch chan<- prometheus.Metric)
}
