package app

import (
	"encoding/json"
	"fmt"

	"github.com/hunterdishner/errors"
)

type Joke struct {
	Joke string `json:"joke"`
}

func (s *AppService) NewJoke() (*Joke, error) {
	resp, err := s.client.Get(s.jokeUrl)
	if err != nil {
		return nil, errors.E(errors.CodeServerError, errors.CodeBadRequest, err)
	}
	defer resp.Body.Close()

	joke := &Joke{}
	err = json.NewDecoder(resp.Body).Decode(joke)
	if err != nil {
		return nil, errors.E(errors.CodeServerError, errors.Decoding, err)
	}

	name, err := s.RandomName()
	if err != nil {
		return nil, err
	}

	joke.Joke = fmt.Sprintf("Hey %s, %s Get it?", name, joke.Joke)

	return joke, nil
}
