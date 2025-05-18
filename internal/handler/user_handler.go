package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/it-gress/itg-go-template/internal/controller"
	"github.com/it-gress/itg-go-template/internal/entities"
	"github.com/it-gress/itg-go-template/internal/utils"
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

// HandleCreateUser handles the creation of a new user.
func (h *UserHandler) HandleCreateUser(c *gin.Context) {
	requestData, err := utils.GetDataFromRequest[entities.CreateUserRequest](c)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	responseData, err := h.UserController.CreateUser(c.Request.Context(), requestData)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User created", responseData)
}

// HandleGetUsers handles the GET request to retrieve all users.
func (h *UserHandler) HandleGetUsers(c *gin.Context) {
	users, err := h.UserController.GetUsers(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Users retrieved", users)
}

// HandleGetUserByID handles the GET request to retrieve a user by ID.
func (h *UserHandler) HandleGetUserByID(c *gin.Context) {
	userID, err := utils.GetParamAsInt(c, "userID")
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	user, err := h.UserController.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User retrieved", user)
}

// HandleUserLogin handles the POST request for user login.
func (h *UserHandler) HandleUserLogin(c *gin.Context) {
	requestData, err := utils.GetDataFromRequest[entities.UserLoginRequest](c)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	responseData, err := h.UserController.UserLogin(c.Request.Context(), requestData)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User logged in", responseData)
}

// HandleUpdateUser handles the PUT request to update a user by ID.
func (h *UserHandler) HandleUpdateUser(c *gin.Context) {
	userID, err := utils.GetParamAsInt(c, "userID")
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	requestData, err := utils.GetDataFromRequest[entities.UpdateUserRequest](c)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	updatedUser, err := h.UserController.UpdateUser(c.Request.Context(), userID, requestData)
	if err != nil {
		utils.ErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User updated", updatedUser)
}
