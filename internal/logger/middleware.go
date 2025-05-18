package logger

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// GinSloggerMiddleware is the middleware that logs request details, userID, scope, and log levels.
func GinSloggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start).Milliseconds()
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userID := c.GetInt("userID")
		scope := c.GetStringSlice("scope")

		level := getLogLevelForStatusCode(statusCode)

		attrs := []slog.Attr{
			slog.Int("status", statusCode),
			slog.String("method", c.Request.Method),
			slog.String("path", path),
			slog.String("ip", clientIP),
			slog.Int64("latency", latency),
			slog.Int("user_id", userID),
			slog.Any("scope", scope),
		}

		slog.LogAttrs(c.Request.Context(), level, "Request processed", attrs...)
	}
}

func getLogLevelForStatusCode(statusCode int) slog.Level {
	if statusCode >= 500 {
		return slog.LevelError
	} else if statusCode >= 400 {
		return slog.LevelWarn
	}

	return slog.LevelInfo
}

// GinDebugPrintRouteFunc is a custom function to print registered routes with slog.
func GinDebugPrintRouteFunc(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	handlerFuncName := handlerName
	if idx := len(handlerName) - 1; idx >= 0 {
		handlerFuncName = handlerName[idx:]
	}

	slog.Debug("Registered route",
		slog.String("method", httpMethod),
		slog.String("path", absolutePath),
		slog.String("handler", handlerFuncName),
		slog.Int("num_handlers", nuHandlers),
	)
}

// GinDebugPrintFunc is a custom debug print function for Gin.
func GinDebugPrintFunc(format string, values ...any) {
	slog.Debug("Debug print",
		slog.String("format", format),
		slog.Any("values", values),
	)
}
