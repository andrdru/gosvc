package metrics

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Namespace = "templateService"
	Subsystem = "backend"

	IsBusinessError func(error) bool
)

func getErrorType(err error) string {
	if err == nil {
		return ""
	}

	if IsBusinessError(err) {
		return "business"
	}

	return "unknown"
}

var (
	hostname, _ = os.Hostname()

	serverHTTP *prometheus.HistogramVec
	clients    *prometheus.HistogramVec
)

func Init() {
	serverHTTP = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "server_http",
			Help:      "http server metrics",
			Buckets:   []float64{.001, .005, .01, .025, .05, .1, .5, 1, 2.5, 5, 10},
		}, []string{"hostname", "method", "error_type"})

	clients = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "clients",
			Help:      "clients metrics",
			Buckets:   []float64{.001, .005, .01, .025, .05, .1, .5, 1, 2.5, 5, 10},
		}, []string{"hostname", "client", "method"})
}

// HistogramObserverServer .
func HistogramObserverServer(method string, err error) prometheus.Observer {
	return serverHTTP.With(map[string]string{
		"hostname":   hostname,
		"method":     method,
		"error_type": getErrorType(err),
	})
}

// HistogramObserverClient .
func HistogramObserverClient(client string, method string) prometheus.Observer {
	return clients.With(map[string]string{
		"hostname": hostname,
		"client":   client,
		"method":   method,
	})
}
