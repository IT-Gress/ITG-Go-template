package controller

import (
	"context"

	"github.com/it-gress/itg-go-template/internal/apierror"
	"github.com/it-gress/itg-go-template/internal/auth"
	"github.com/it-gress/itg-go-template/internal/entities"
	"github.com/it-gress/itg-go-template/internal/models"
	"github.com/it-gress/itg-go-template/internal/repository"
)

// UserController is a struct that holds the user repository.
type UserController struct {
	UserRepository *repository.UserRepository
}

// NewUserController initializes a new UserController with the given user repository.
func NewUserController(userRepository *repository.UserRepository) *UserController {
	return &UserController{
		UserRepository: userRepository,
	}
}

// CreateUser creates a new user in the repository and returns the created user.
func (uc *UserController) CreateUser(
	c context.Context, createUserRequest *entities.CreateUserRequest) (*entities.UserDTO, *apierror.APIError) {
	// Convert CreateUserRequest to User entity
	newUser := &models.User{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Username: createUserRequest.Username,
		RoleID:   createUserRequest.RoleID,
	}

	hash, err := auth.CreateHash(createUserRequest.Password)
	if err != nil {
		return nil, err
	}

	newUser.PasswordHash = hash

	// Save the new user to the repository
	savedUser, err := uc.UserRepository.InsertUser(c, newUser)
	if err != nil {
		return nil, err
	}

	return savedUser.ToDTO(), nil
}

// GetUsers retrieves all users from the repository and converts them to UserDTOs.
func (uc *UserController) GetUsers(c context.Context) ([]*entities.UserDTO, *apierror.APIError) {
	users, err := uc.UserRepository.FindAllUsers(c)
	if err != nil {
		return nil, err
	}

	userDTOs := make([]*entities.UserDTO, len(users))

	for i, user := range users {
		userDTOs[i] = user.ToDTO()
	}

	return userDTOs, nil
}

// GetUserByID retrieves a user by their ID from the repository and converts it to a UserDTO.
func (uc *UserController) GetUserByID(c context.Context, userID int) (*entities.UserDTO, *apierror.APIError) {
	user, err := uc.UserRepository.FindUserByID(c, userID)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

// UpdateUser updates an existing user in the repository and returns the updated user.
func (uc *UserController) UpdateUser(
	c context.Context,
	userID int,
	updateUserRequest *entities.UpdateUserRequest) (*entities.UserDTO, *apierror.APIError) {
	// Find the user by ID
	user, err := uc.UserRepository.FindUserByID(c, userID)
	if err != nil {
		return nil, err
	}

	// Update the user fields
	user.Name = updateUserRequest.Name
	user.Email = updateUserRequest.Email
	user.Username = updateUserRequest.Username
	user.RoleID = updateUserRequest.RoleID
	user.IsActive = updateUserRequest.IsActive

	if *updateUserRequest.Password != "" && updateUserRequest.Password != nil {
		hash, err := auth.CreateHash(*updateUserRequest.Password)
		if err != nil {
			return nil, err
		}

		user.PasswordHash = hash
	}

	// Save the updated user to the repository
	updatedUser, err := uc.UserRepository.UpdateUser(c, user)
	if err != nil {
		return nil, err
	}

	return updatedUser.ToDTO(), nil
}
