package server

import (
	"fmt"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/it-gress/itg-go-template/internal/config"
	"github.com/it-gress/itg-go-template/internal/handler"
)

// Server is a struct that holds the configuration, handlers, and router for the server.
type Server struct {
	cfg      *config.Config
	handlers *handler.Handlers
	router   *gin.Engine
}

// NewServer initializes a new Server with the given configuration and handlers.
func NewServer(cfg *config.Config, handlers *handler.Handlers) *Server {
	gin.DebugPrintRouteFunc = ginDebugPrintRouteFunc
	gin.DebugPrintFunc = ginDebugPrintFunc

	if cfg.Environment == "production" {
		slog.Info("Running gin in production mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		slog.Info("Running gin in development mode")
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.Use(ginSloggerMiddleware())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	server := &Server{
		cfg:      cfg,
		handlers: handlers,
		router:   r,
	}

	return server
}

// Start starts listening on the configured port.
func (s *Server) Start() error {
	if err := s.router.Run(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	slog.Info("Server started successfully", slog.Int("port", s.cfg.Port))

	return nil
}

// ginDebugPrintRouteFunc is a custom function to print registered routes with slog.
func ginDebugPrintRouteFunc(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	slog.Debug("Registered route",
		slog.String("method", httpMethod),
		slog.String("path", absolutePath),
		slog.String("handler", handlerName),
		slog.Int("num_handlers", nuHandlers),
	)
}

// ginDebugPrintFunc is a custom debug print function for Gin.
func ginDebugPrintFunc(format string, values ...any) {
	slog.Debug("Debug print",
		slog.String("format", format),
		slog.Any("values", values),
	)
}
