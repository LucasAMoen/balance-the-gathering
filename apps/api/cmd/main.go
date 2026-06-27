package main

import (
	"log/slog"
	"os"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "localhost"
	}

	cfg := config{
		address:  serverAddress,
		port:     "8040",
		database: dbconfig{},
	}

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	if error := api.run(api.mount()); error != nil {
		slog.Error("Server has failed to start", "error", error)
		os.Exit(1)
	}
}
