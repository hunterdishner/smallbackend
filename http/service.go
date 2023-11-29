package http

import (
	"context"

	"github.com/hunterdishner/gomux"
	"github.com/hunterdishner/wildfire/app"
	"github.com/rs/cors"
)

type HttpService struct {
	router *gomux.Server
	app    *app.AppService
	// whatever else, message queue system, cache, websockets, etc
}

func New(app *app.AppService) *HttpService {
	s := &HttpService{
		app: app,
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://localhost"}, //swap to configurable variable with JSON file or something
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
	})

	s.router = gomux.New(context.Background(), "api", gomux.Port(10000), gomux.CustomCors(cors)) //needs TLS cert but using option gomux.TLS() and naming the files appropriately

	s.router.AddRoutes(
		gomux.Get("/joke", Joke),
	)

	//add panic recovery or some sort of log fatal
	return s
}

func (s *HttpService) Serve() error {
	return s.router.Serve()
}
