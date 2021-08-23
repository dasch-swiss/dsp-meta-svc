package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

// service implements Service interface
type service struct {
	pHistogram           *prometheus.HistogramVec
	httpRequestHistogram *prometheus.HistogramVec
}

// NewPrometheusService create a new prometheus service
func NewPrometheusService() (*service, error) {
	cli := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "pushgateway",
		Name:      "cmd_duration_seconds",
		Help:      "CLI application execution in seconds",
		Buckets:   prometheus.DefBuckets,
	}, []string{"name"})
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &service{
		pHistogram:           cli,
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.pHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	err = prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

// SaveHTTP send metrics to server
func (s *service) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
