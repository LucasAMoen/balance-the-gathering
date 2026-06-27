package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *application) mount() http.Handler {
	router := gin.Default()

	// Config
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://" + app.config.address + ":5173"}
	config.AllowHeaders = []string{"Accept", "Access-Control-Allow-Origin", "Referer", "sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "user-agent"}

	// Middleware
	router.Use(cors.New(config))
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Routes
	router.GET("/health", getHealth)
	router.GET("/cards", getCards)
	router.GET("/card", getCard)

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
