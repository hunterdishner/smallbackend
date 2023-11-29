package app

import (
	"encoding/json"
	"fmt"

	"github.com/hunterdishner/errors"
)

type RandomName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s *AppService) RandomName() (string, error) {
	resp, err := s.client.Get(s.nameUrl)
	if err != nil {
		return "", errors.E(errors.CodeServerError, errors.CodeBadRequest, err)
	}
	defer resp.Body.Close()

	name := &RandomName{}
	err = json.NewDecoder(resp.Body).Decode(name)
	if err != nil {
		return "", errors.E(errors.CodeServerError, errors.Decoding, err)
	}

	return fmt.Sprintf("%s %s", name.FirstName, name.LastName), nil
}
