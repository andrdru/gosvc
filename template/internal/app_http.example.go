package internal

import (
	"log"
	"net/http"

	"github.com/andrdru/gosvc/apimsg"
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
	var msg = apimsg.NewJson()

	type response struct {
		Hello     string `json:"hello"`
		ParamGET  string `json:"param_get"`
		ParamPOST string `json:"param_post"`
	}

	var resp = &response{
		Hello: "hello world!",
	}

	// read request param
	var p = r.URL.Query().Get("param")
	if p != "" {
		resp.ParamGET = p
	}

	// read request body
	var body = m.GetBodyCtx(r.Context())
	if len(body) > 0 {
		resp.ParamPOST = string(body)
	}

	msg.Data = resp

	var err = msg.Return(w)
	if err != nil {
		log.Printf("write response: %s", err)
	}
}
