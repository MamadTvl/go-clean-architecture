package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	IncrementUserCreation()
}

type PrometheusMetrics struct {
	userCreationCounter prometheus.Counter
}

func NewPrometheusMetrics() Metrics {
	var metrics Metrics
	userCreationCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "created_users_total",
			Help: "Total number of users created successfully",
		},
	)
	prometheus.MustRegister(userCreationCounter)
	metrics = &PrometheusMetrics{
		userCreationCounter,
	}
	return metrics
}

func (p *PrometheusMetrics) IncrementUserCreation() {
	p.userCreationCounter.Inc()
}
