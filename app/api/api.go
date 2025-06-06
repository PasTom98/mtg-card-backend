package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Address string
	Version string
}

type Application struct {
	Config Config
}

func (api *Application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	//set a timeout value on the request context (ctx), that will signal
	//through ctx.Done() that the request hast timed out and further processing should be stopped
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	//health + stats
	r.Route("/"+api.Config.Version, func(r chi.Router) {
		r.Get("/health", api.HealthCheckHandler)
	})

	//cards

	//users

	//auth

	return r
}

func (api *Application) Run(handler http.Handler) error {

	srv := &http.Server{
		Addr:         api.Config.Address,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	serverAddress := "http://localhost" + api.Config.Address + "/" + api.Config.Version + "/health"
	log.Printf("Starting server at %s", serverAddress)

	return srv.ListenAndServe()
}
