package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/LucasAMoen/balance-the-gathering/application/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	serverAddress := env.GetEnvVariable("SERVER_ADDRESS", "localhost")
	dbConnectionString := env.GetEnvVariable("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=balance_the_gathering sslmode=verify-full")

	cfg := config{
		address: serverAddress,
		port:    "8040",
		database: dbconfig{
			connectionString: dbConnectionString,
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	conn, error := pgx.Connect(ctx, cfg.database.connectionString)
	if error != nil {
		panic(error)
	}
	defer conn.Close(ctx)
	logger.Info("Connected to database")

	api := application{
		config:   cfg,
		database: conn,
	}

	if error := api.run(api.mount()); error != nil {
		slog.Error("Server has failed to start", "error", error)
		os.Exit(1)
	}
}
