package middlewares

import (
	"net/http"
	"strings"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	var wrapped = h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func isContentTypeMultipart(r *http.Request) bool {
	var contentType = r.Header.Get("Content-Type")
	if contentType == "" {
		return false
	}

	contentType = strings.ToLower(strings.TrimSpace(contentType))
	return strings.HasPrefix(contentType, "multipart/form-data")
}
