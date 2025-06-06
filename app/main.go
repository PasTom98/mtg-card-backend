package main

import (
	"log"
	"mtg-card-backend/app/api"
)

func main() {
	cfg := api.Config{
		Address: ":8080",
		Version: "0.0.1",
	}

	app := &api.Application{
		Config: cfg,
	}

	log.Fatal(app.Run())
}
