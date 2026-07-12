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

	cfg := Config{
		Address: serverAddress,
		Port:    "8040",
		Database: Dbconfig{
			ConnectionString: dbConnectionString,
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	conn, error := pgx.Connect(ctx, cfg.Database.ConnectionString)
	if error != nil {
		panic(error)
	}
	defer conn.Close(ctx)
	logger.Info("Connected to database")

	app := Application{
		Config:   cfg,
		Database: conn,
	}

	if error := app.Run(app.Mount()); error != nil {
		slog.Error("Server has failed to start", "error", error)
		os.Exit(1)
	}
}
