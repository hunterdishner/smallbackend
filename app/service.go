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
	//TODO: database connection service goes here
	//TODO: roles/permissions service gets added here
	//Typically I write the basic validations of JWT authenticity into the reverse proxy server (or something like Auth0 does it for us). We just extract the individual service permissions in the web service level.
}

func New(cfg *config.Config) *AppService {
	c := &http.Client{Timeout: 10 * time.Second} //ToDo: arbitrary timeout, needs refined for this service

	return &AppService{client: c, nameUrl: cfg.NameApiURL, jokeUrl: cfg.JokeApiURL} //normally has things like the auth service instantiated to validate specific permissions (seperate from middleware in reverse proxy), database connections, mail clients, etc
}
