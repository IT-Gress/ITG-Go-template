package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes the routes for the server.
func (s *Server) RegisterRoutes() {
	v1 := s.router.Group("/api/v1")

	auth := v1.Group("/auth")
	// /api/v1/auth
	auth.POST("/login", s.handlers.UserHandler.HandleUserLogin)

	users := v1.Group("/users", s.requireAuthentication)
	// /api/v1/users
	users.POST("/", s.requirePermissionsOrOwnResource("users.create"), s.handlers.UserHandler.HandleCreateUser)
	users.GET("/", s.requirePermissionsOrOwnResource("users.read"), s.handlers.UserHandler.HandleGetUsers)
	users.GET("/:userID", s.requirePermissionsOrOwnResource("users.read"), s.handlers.UserHandler.HandleGetUserByID)
	users.PUT("/:userID", s.requirePermissionsOrOwnResource("users.update"), s.handlers.UserHandler.HandleUpdateUser)

	s.router.GET("/health", s.healthHandler)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}
