package middlewares

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog"
)

type (
	bodyCtxKey string
)

var key bodyCtxKey = "body"

// SetBodyCtx add request body to context
// do not read it twice
func SetBodyCtx(ctx context.Context, bytes []byte) context.Context {
	return context.WithValue(ctx, key, bytes)
}

// GetBodyCtx get request body from context
func GetBodyCtx(ctx context.Context) (body []byte) {
	var value = ctx.Value(key)
	if value == nil {
		return nil
	}
	return value.([]byte)
}

// BodyReader read request body to ctx
var BodyReader = func(logger *zerolog.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if isContentTypeMultipart(r) {
				next.ServeHTTP(w, r)
				return
			}

			var bytes, err = ioutil.ReadAll(r.Body)
			if err != nil {
				logger.Err(err).Msg("could not read request body")
				return
			}

			r = r.WithContext(SetBodyCtx(r.Context(), bytes))
			next.ServeHTTP(w, r)
		}
	}
}
