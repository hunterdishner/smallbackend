package app

import (
	"encoding/json"
	"net/http"

	"github.com/hunterdishner/errors"
)

type RandomName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s *AppService) RandomName() (*RandomName, error) {
	resp, err := s.client.Get(s.nameUrl)
	if err != nil {
		return nil, errors.E(errors.CodeServerError, errors.CodeBadRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		name := &RandomName{}
		err = json.NewDecoder(resp.Body).Decode(name)
		if err != nil {
			return nil, errors.E(errors.CodeServerError, errors.Decoding, err)
		}

		return name, nil
	}

	return &RandomName{FirstName: "John", LastName: "Doe"}, nil //This service rate limits at about 10 request per second, so lets give a default if we hit that.
}
