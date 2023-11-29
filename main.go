package main

import (
	"log"

	"github.com/hunterdishner/wildfire/app"
	"github.com/hunterdishner/wildfire/config"
	httpd "github.com/hunterdishner/wildfire/http"
	"github.com/jinzhu/configor"
)

func main() {
	cfg := &config.Config{}
	err := configor.Load(cfg, "config.json")
	if err != nil {
		log.Fatalf("Error loading config file, %+v", err)
		return // lets not run unconfigured
	}
	//custom cors as default allows all origins for testing
	app := app.New()

	server := httpd.New(app, cfg)

	log.Fatal(server.Serve())
}
