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
	cfg    *config.Config
	// whatever else, message queue system, cache, websockets, etc
}

func New(app *app.AppService, config *config.Config) *HttpService {
	s := &HttpService{
		app: app,
		cfg: config,
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{config.AllowedOrigins},
		AllowCredentials: config.CorsCredentials,
		AllowedMethods:   config.AllowedMethods,
		AllowedHeaders:   config.AllowedHeaders,
	})

	s.router = gomux.New(context.Background(), "api", gomux.Port(10000), gomux.CustomCors(cors)) //TODO: needs TLS cert but using option gomux.TLS() and naming the files server.crt and server.key.
	//I didn't generate these for this repo due to time. Test Url goes to https once that option is passed in.

	s.router.AddRoutes(
		gomux.Get("/joke", s.Joke),
	)

	return s
}

func (s *HttpService) Serve() error {
	return s.router.Serve()
}
