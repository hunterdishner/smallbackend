package app

import (
	"encoding/json"
	"fmt"

	"github.com/hunterdishner/errors"
)

type Joke struct {
	Type  string `json:"type"`
	Value struct {
		Categories []string `json:"categories"`
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
	} `json:"value"`
}

func (s *AppService) NewJoke(firstName string, lastName string) (string, error) {

	if firstName == "" || lastName == "" { //option to pass in your own name just incase
		name, err := s.RandomName()
		if err != nil {
			return "", err
		}

		firstName = name.FirstName
		lastName = name.LastName
	}

	resp, err := s.client.Get(fmt.Sprintf("%s&firstName=%s&lastName=%s", s.jokeUrl, firstName, lastName))
	if err != nil {
		return "", errors.E(errors.CodeServerError, errors.CodeBadRequest, err)
	}
	defer resp.Body.Close()

	joke := &Joke{}
	err = json.NewDecoder(resp.Body).Decode(joke)
	if err != nil {
		return "", errors.E(errors.CodeServerError, errors.Decoding, err)
	}

	return joke.Value.Joke, nil
}
