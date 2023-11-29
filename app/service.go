package app

import (
	"net/http"
	"time"

	"github.com/hunterdishner/wildfire/config"
)

type AppService struct {
	client  *http.Client // client for api calls
	nameUrl string
	jokeUrl string
	//dbcon goes here
	//roles/permissions service gets added here
	//etc
}

func New(cfg *config.Config) *AppService {
	c := &http.Client{Timeout: 10 * time.Second} //arbitrary timeout, needs refined for this service

	return &AppService{client: c, nameUrl: cfg.NameApiURL, jokeUrl: cfg.JokeApiURL} //normally has things like the auth service instantiated to validate specific permissions (seperate from middleware), database connections, mail clients, etc
}
