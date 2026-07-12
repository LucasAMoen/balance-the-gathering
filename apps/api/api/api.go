package main

import (
	"log"
	"net/http"
	"time"

	"github.com/LucasAMoen/balance-the-gathering/application/cards"
	"github.com/LucasAMoen/balance-the-gathering/infrastructure"
	repository "github.com/LucasAMoen/balance-the-gathering/infrastructure/adapters/postgresql/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
)

func (app *Application) Mount() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://" + app.Config.Address + ":5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Access-Control-Allow-Origin", "Referer", "sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "user-agent"},
	}))

	// Infrastructure
	repo := infrastructure.NewRepository(repository.New(app.Database))

	// Handlers
	cardHandler := cards.NewHandler(repo)

	// Routes
	router.Get("/health", cardHandler.GetHealth)
	router.Get("/cards", cardHandler.GetCards)
	router.Get("/card", cardHandler.GetCard)

	return router
}

func (app *Application) Run(handler http.Handler) error {
	server := &http.Server{
		Addr:         app.Config.Address + ":" + app.Config.Port,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at address: %s:%s", app.Config.Address, app.Config.Port)

	return server.ListenAndServe()
}

type Application struct {
	Config   Config
	Database *pgx.Conn
}

type Config struct {
	Address  string
	Port     string
	Database Dbconfig
}

type Dbconfig struct {
	ConnectionString string
}
