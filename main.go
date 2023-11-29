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

	app := app.New() //business logic layer

	server := httpd.New(app, cfg)

	log.Fatal(server.Serve()) //may need panic recovery depending on what environment its hosted on.
	//GCP had self healing so we cared more about logging the crash since GCP would self heal the container
}
