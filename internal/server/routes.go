package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes the routes for the server.
func (s *Server) RegisterRoutes() {
	s.router.GET("/health", s.healthHandler)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}
