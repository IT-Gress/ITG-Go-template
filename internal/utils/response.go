package utils

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/it-gress/itg-go-template/internal/apierror"
)

// SuccessResponse sends a JSON response with the given status code, message, and data.
func SuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

// ErrorResponse handles errors and sends a JSON response with the appropriate status code and message.
func ErrorResponse(c *gin.Context, err error) {
	var apiErr *apierror.APIError
	if errors.As(err, &apiErr) {
		if apiErr != nil {
			slog.Warn("APIError", slog.Any("error", apiErr.Err), slog.String("message", apiErr.Message))
		}

		c.JSON(apiErr.Code, gin.H{
			"message": apiErr.Message,
		})

		return
	}

	slog.Error("Unknown ErrorResponse", slog.Any("error", err))
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": http.StatusText(http.StatusInternalServerError),
	})
}
