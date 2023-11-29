package http

import (
	"io"
	"net/http"

	"github.com/hunterdishner/errors"
)

func Joke(w io.Writer, r *http.Request) (interface{}, error) {
	return nil, errors.E(errors.CodeServerError, errors.Canceled, "Route not setup yet")
}
