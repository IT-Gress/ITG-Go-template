package server

import (
	"log/slog"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/it-gress/itg-go-template/internal/apierror"
	"github.com/it-gress/itg-go-template/internal/auth"
	"github.com/it-gress/itg-go-template/internal/utils"
)

const authHeaderKey = "Authorization"
const localsUserIDKey = "userID"
const localsScopesKey = "scopes"

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

	c.Set(localsScopesKey, claims.Scopes)

	userID, convertErr := strconv.Atoi(claims.Subject)
	if convertErr != nil {
		utils.ErrorResponse(c, apierror.New(http.StatusUnauthorized, "Invalid JWT Subject", convertErr))
		c.Abort()

		return
	}

	c.Set(localsUserIDKey, userID)

	c.Next()
}

func (s *Server) requirePermissionsOrOwnResource(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user owns the resource
		userID, _ := utils.GetParamAsInt(c, "userID")
		if userID == c.GetInt(localsUserIDKey) && userID != 0 {
			c.Next()
			return
		}

		// Get the scopes from the context
		scopes := c.GetStringSlice(localsScopesKey)

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
	authHeader := c.GetHeader(authHeaderKey)
	if authHeader == "" {
		return ""
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token
}

// ginSloggerMiddleware is the middleware that logs request details, userID, scope, and log levels.
func ginSloggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start).Milliseconds()
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userID := c.GetInt(localsUserIDKey)
		scopes := c.GetStringSlice(localsScopesKey)

		level := getLogLevelForStatusCode(statusCode)

		attrs := []slog.Attr{
			slog.Int("status", statusCode),
			slog.String("method", c.Request.Method),
			slog.String("path", path),
			slog.String("ip", clientIP),
			slog.Int64("latency", latency),
			slog.Int("user_id", userID),
			slog.Any("scope", scopes),
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
