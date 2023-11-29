package http

import (
	"context"

	"github.com/hunterdishner/gomux"
	"github.com/hunterdishner/wildfire/app"
	"github.com/hunterdishner/wildfire/config"
	"github.com/rs/cors"
)

type HttpService struct {
	router *gomux.Server
	app    *app.AppService
	// whatever else, message queue system, cache, websockets, etc
}

func New(app *app.AppService, config *config.Config) *HttpService {
	s := &HttpService{
		app: app,
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{config.AllowedOrigins},
		AllowCredentials: config.CorsCredentials,
		AllowedMethods:   config.AllowedMethods,
		AllowedHeaders:   config.AllowedHeaders,
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
