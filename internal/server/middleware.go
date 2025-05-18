package server

import (
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/it-gress/itg-go-template/internal/apierror"
	"github.com/it-gress/itg-go-template/internal/auth"
	"github.com/it-gress/itg-go-template/internal/utils"
)

func (s *Server) requireAuthentication(c *gin.Context) {
	// Get jwt from request haeder
	token := extractToken(c)
	if token == "" {
		utils.ErrorResponse(c, apierror.New(http.StatusUnauthorized, "No token provided", nil))
		c.Abort()

		return
	}

	// Check if token is valid
	claims, err := auth.ValidateToken(token)

	if err != nil {
		utils.ErrorResponse(c, apierror.New(http.StatusUnauthorized, "Invalid token", err))
		c.Abort()

		return
	}

	c.Set("scopes", claims.Scopes)

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		utils.ErrorResponse(c, apierror.New(http.StatusUnauthorized, "Invalid JWT Subject", err))
		c.Abort()

		return
	}

	c.Set("userID", userID)

	c.Next()
}

func (s *Server) requirePermissionsOrOwnResource(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user owns the resource
		userID, _ := utils.GetParamAsInt(c, "userID")
		if userID == c.GetInt("userID") && userID != 0 {
			c.Next()
			return
		}

		// Get the scopes from the context
		scopes := c.GetStringSlice("scopes")

		// Check if the required permission is in auth scopes
		if !slices.Contains(scopes, permission) {
			utils.ErrorResponse(c, apierror.New(http.StatusForbidden, http.StatusText(http.StatusForbidden), nil))
			c.Abort()

			return
		}

		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token
}
