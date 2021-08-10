package middleware

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/metric"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
)

// Metrics to prometheus
func Metrics(metricService metric.Service) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		appMetric := metric.NewHTTP(r.URL.Path, r.Method)
		appMetric.Started()
		next(w, r)
		res := w.(negroni.ResponseWriter)
		appMetric.Finished()
		appMetric.StatusCode = strconv.Itoa(res.Status())
		metricService.SaveHTTP(appMetric)
	}
}
