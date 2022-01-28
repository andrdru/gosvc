package internal

import (
	"net/http"

	m "github.com/andrdru/gosvc/middlewares"
	"github.com/rs/zerolog"

	"github.com/andrdru/gosvc/template/middlewares"
)

func init() {
	// redefine metrics port
	metricsAddr = ":1234"

	// redefine app initial handler
	runHandler = func(fv flags, logger *zerolog.Logger) {
		var mw = []m.Middleware{
			m.BodyReader(logger),
			m.Logger(logger),
			middlewares.HTTPServerLatency(logger),
		}

		var mux = http.NewServeMux()

		mux.Handle("/", m.MultipleMiddleware(handlerExample, mw...))
		_ = http.ListenAndServe(":8080", mux)
	}
}

func handlerExample(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))

	// read request param
	var p = r.URL.Query().Get("param")
	if p != "" {
		_, _ = w.Write([]byte(" "))
		_, _ = w.Write([]byte(p))
	}

	// read request body
	var body = m.GetBodyCtx(r.Context())
	if len(body) > 0 {
		_, _ = w.Write([]byte(" "))
		_, _ = w.Write(body)
	}
}
