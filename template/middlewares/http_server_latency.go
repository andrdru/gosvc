package middlewares

import (
	"net/http"
	"time"

	"github.com/andrdru/gosvc/middlewares"
	"github.com/rs/zerolog"

	"github.com/andrdru/gosvc/template/metrics"
)

// HTTPServerLatency .
var HTTPServerLatency = func(logger *zerolog.Logger) middlewares.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var methodPath = r.Method + " " + r.URL.Path
			var startedAt = time.Now()

			next.ServeHTTP(w, r)

			// todo pass err to metrics
			metrics.HistogramObserverServer(methodPath, nil).Observe(time.Since(startedAt).Seconds())
		}
	}
}
