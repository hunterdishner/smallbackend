package http

import (
	"io"
	"net/http"
)

func (s *HttpService) Joke(w io.Writer, r *http.Request) (interface{}, error) {
	values := r.URL.Query()

	return s.app.NewJoke(values.Get("firstName"), values.Get("lastName")) // muxxer handles the error
}
