package config

import (
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds the application configuration.
type Config struct {
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
	Port        int    `envconfig:"PORT" default:"8080"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	Secret      string `envconfig:"SECRET" default:"thisisasecret"`
	HashSalt    string `envconfig:"HASH_SALT" default:"thisisahashsalt"`
	Database    Database
}

// Database holds the configuration for the PostgreSQL database connection.
type Database struct {
	User     string `envconfig:"POSTGRES_USERNAME"  default:"postgres"`
	Password string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port     int    `envconfig:"POSTGRES_PORT" default:"5432"`
	Database string `envconfig:"POSTGRES_DATABASE" default:"basecrm"`
}

// PostgresDSN generates the PostgreSQL Data Source Name (DSN) for connecting to the database.
func (c Database) PostgresDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable client_encoding=UTF8",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}

// LoadConfiguration loads the configuration from environment variables and .env file.
func LoadConfiguration() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Info("no .env file found, using system environment variables")
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
