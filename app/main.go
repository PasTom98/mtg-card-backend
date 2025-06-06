package main

import (
	"log"
	"main/api"
)

func main() {
	cfg := api.Config{
		Address: ":8080",
		Version: "0.0.1",
	}

	app := &api.Application{
		Config: cfg,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
