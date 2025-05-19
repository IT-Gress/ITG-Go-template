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

	ensureDefaultUser(db)

	if err := server.Start(); err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		panic(err)
	}
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
	// Check if there is atleast one user in the database
	var count int

	err := db.Get(&count, "SELECT COUNT(*) FROM users")
	if err != nil {
		slog.Error("failed to check user count", slog.Any("error", err))
		panic(err)
	}

	if count > 0 {
		slog.Debug("a user already exists, skipping default user creation")
		return
	}

	password := "admin"
	username := "admin"

	passwordHash, hashErr := auth.CreateHash(password)
	if hashErr != nil {
		slog.Error("failed to create password hash", slog.Any("error", hashErr))
		panic(hashErr)
	}

	// Insert a default user
	_, err = db.Exec("INSERT INTO users (name, username, password_hash, email, role_id) VALUES ($1, $2, $3, $4, $5)",
		"Administrator", username, passwordHash, "admin@example.com", 1)
	if err != nil {
		slog.Error("failed to create default user", slog.Any("error", err))
		panic(err)
	}

	slog.Info("default user created successfully", slog.String("username", username), slog.String("password", password))
}
