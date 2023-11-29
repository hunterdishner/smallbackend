package http

import (
	"io"
	"net/http"

	"github.com/hunterdishner/errors"
)

func (s *HttpService) Joke(w io.Writer, r *http.Request) (interface{}, error) {
	values := r.URL.Query()

	if values.Get("forceError") == "1" {
		return nil, errors.E(errors.CodeServerError, errors.Decoding, "Forced Error for demonstration purposes")
	}

	return s.app.NewJoke(values.Get("firstName"), values.Get("lastName")) // muxxer handles the error
}
