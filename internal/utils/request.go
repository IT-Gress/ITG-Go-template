package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/it-gress/itg-go-template/internal/apierror"
)

// GetDataFromRequest binds the incoming request data to a struct of type T.
// It returns the struct and any binding error encountered.
func GetDataFromRequest[T any](c *gin.Context) (*T, *apierror.APIError) {
	data := new(T)

	err := c.ShouldBind(data)
	if err != nil {
		return nil, apierror.New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	return data, nil
}

// GetParamAsInt retrieves a URL parameter by name and converts it to an integer.
func GetParamAsInt(c *gin.Context, name string) (int, *apierror.APIError) {
	param, err := GetParam(c, name)
	if err != nil {
		return 0, err
	}

	value, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return 0, apierror.New(http.StatusBadRequest, "Invalid parameter", convertErr)
	}

	return value, nil
}

// GetParam retrieves a URL parameter by name.
func GetParam(c *gin.Context, name string) (string, *apierror.APIError) {
	param := c.Param(name)
	if param == "" {
		return "", apierror.New(http.StatusBadRequest, "Missing parameter", nil)
	}

	return param, nil
}
