package main

import (
	"log"
	"net/http"
	"time"

	"github.com/LucasAMoen/balance-the-gathering/application/cards"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://" + app.config.address + ":5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Access-Control-Allow-Origin", "Referer", "sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "user-agent"},
	}))

	// Services
	cardService := cards.NewService()
	// Handlers
	cardHandler := cards.NewHandler(cardService)

	// Routes
	router.Get("/health", cardHandler.GetHealth)
	router.Get("/cards", cardHandler.GetCards)
	router.Get("/card", cardHandler.GetCard)

	return router
}

func (app *application) run(handler http.Handler) error {
	server := &http.Server{
		Addr:         app.config.address + ":" + app.config.port,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at address: %s:%s", app.config.address, app.config.port)

	return server.ListenAndServe()
}

type application struct {
	config config
}

type config struct {
	address  string
	port     string
	database dbconfig
}

type dbconfig struct {
	dsn string
}
