package main

import (
	"log"

	"github.com/hunterdishner/wildfire/app"
	httpd "github.com/hunterdishner/wildfire/http"
)

func main() {

	//custom cors as default allows all origins for testing
	app := app.New()

	server := httpd.New(app)

	log.Fatal(server.Serve())
}
