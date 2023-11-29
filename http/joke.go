package http

import (
	"io"
	"net/http"
)

func (s *HttpService) Joke(w io.Writer, r *http.Request) (interface{}, error) {
	return s.app.NewJoke() // muxxer handles the error
}
