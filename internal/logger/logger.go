package logger

import (
	"log/slog"
	"os"
)

// Init initializes the logger with the specified log level and sets it as default.
func Init(logLevel string) {
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(os.Stdout,
				&slog.HandlerOptions{
					Level: parseLevel(logLevel),
				},
			),
		),
	)
}

func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "info", "":
		return slog.LevelInfo
	default:
		return slog.LevelInfo
	}
}
