package main

import (
	"log/slog"

	"github.com/it-gress/itg-go-template/internal/auth"
	"github.com/it-gress/itg-go-template/internal/config"
	"github.com/it-gress/itg-go-template/internal/controller"
	"github.com/it-gress/itg-go-template/internal/database"
	"github.com/it-gress/itg-go-template/internal/handler"
	"github.com/it-gress/itg-go-template/internal/logger"
	"github.com/it-gress/itg-go-template/internal/repository"
	"github.com/it-gress/itg-go-template/internal/server"
	"github.com/jmoiron/sqlx"
)

func init() {
	logger.Init("debug")
}

func main() {
	cfg := mustLoadConfig()
	logger.Init(cfg.LogLevel)

	auth.Init(cfg.Secret, cfg.HashSalt)
	db := mustInitDatabase(&cfg.Database)

	repos := repository.NewRepositories(db)
	controllers := controller.NewControllers(repos)
	handlers := handler.NewHandlers(controllers)

	server := server.NewServer(cfg, handlers)
	server.RegisterRoutes()

	if err := server.Start(); err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		panic(err)
	}

	ensureDefaultUser(db)
}

func mustLoadConfig() *config.Config {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		slog.Error("failed to load configuration", slog.Any("error", err))
		panic(err)
	}

	slog.Info("loaded configuration", slog.Any("config", cfg))

	return cfg
}

func mustInitDatabase(cfg *config.Database) *sqlx.DB {
	dsn := cfg.PostgresDSN()
	slog.Debug("initializing database connection", slog.String("dsn", dsn))

	db, err := database.Init(dsn)
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("error", err))
		panic(err)
	}

	slog.Info("Database connection / migration successful")

	return db
}

func ensureDefaultUser(db *sqlx.DB) {

}
