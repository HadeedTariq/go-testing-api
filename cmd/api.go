package main

import (
	"log"
	"net/http"
	"time"

	"github.com/HadeedTariq/go-testing-api/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	addr     string
	dbConfig dbConfig
}

type dbConfig struct {
	dsn string
}

type application struct {
	config config
}

func (app *application) mount() http.Handler {
	// ~ so over there the router related stuff is going to manage over there
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	// ~ so these are the middlewares which are by default will be use over there for checking the request related stuff
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("All good"))
	})

	productsHandler := products.NewHandler(nil)
	r.Get("/products", productsHandler.ListProducts)

	return r
}

func (app *application) run(h http.Handler) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at addr %s", app.config.addr)

	return server.ListenAndServe()
}
