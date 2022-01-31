package middlewares

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog"
)

var LoggerHandler = func(r *http.Request, logger *zerolog.Logger) {
	var l = logger.Info()
	if (r.Method == http.MethodPost || r.Method == http.MethodPut) &&
		!isContentTypeMultipart(r) {
		var b, err = ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Err(err).Msg("could not read request body")
			return
		}
		if b != nil {
			l = l.Bytes("body", b)
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}

	l.Msgf("HTTP request: %s %s", r.Method, r.RequestURI)
}

var Logger = func(logger *zerolog.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			LoggerHandler(r, logger)
			next.ServeHTTP(w, r)
		}
	}
}
