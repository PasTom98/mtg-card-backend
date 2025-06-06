package api

import (
	"log"
	"net/http"
	"time"
)

type Config struct {
	Address string
	Version string
}

type Application struct {
	Config Config
}

func (api *Application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /"+api.Config.Version+"/health", api.HealthCheckHandler)

	return mux
}

func (api *Application) Run() error {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         api.Config.Address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	serverAddress := "http://localhost" + api.Config.Address + "/"
	log.Printf("Starting server at %s", serverAddress)

	return srv.ListenAndServe()
}
