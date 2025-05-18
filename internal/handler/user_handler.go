package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/it-gress/itg-go-template/internal/controller"
)

// UserHandler is a struct that holds the user controller.
type UserHandler struct {
	UserController *controller.UserController
}

// NewUserHandler initializes a new UserHandler with the given user controller.
func NewUserHandler(userController *controller.UserController) *UserHandler {
	return &UserHandler{
		UserController: userController,
	}
}

// HandleGetUsers handles the GET request to retrieve all users.
func (h *UserHandler) HandleGetUsers(c *gin.Context) {
	users, err := h.UserController.GetUsers(c.Request.Context())
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
