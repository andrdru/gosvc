package apimsg

import (
	"net/http"
)

type (
	Message interface {
		Return(w http.ResponseWriter) error
		SetError(options ...Option)
		Raw() interface{}
	}
)
