package app

import (
	"net/http"
	"time"
)

type AppService struct {
	client *http.Client // client for api calls
	//dbcon goes here
	//roles/permissions service gets added here
	//etc
}

func New() *AppService {
	c := &http.Client{Timeout: 10 * time.Second} //arbitrary timeout, needs refined for this service

	return &AppService{client: c} //normally has things like the auth service instantiated to validate specific permissions (seperate from middleware), database connections, mail clients, etc
}
