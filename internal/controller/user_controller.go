package controller

import (
	"context"

	"github.com/it-gress/itg-go-template/entities"
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

// GetUsers retrieves all users from the repository and converts them to UserDTOs.
func (uc *UserController) GetUsers(c context.Context) ([]*entities.UserDTO, error) {
	users, err := uc.UserRepository.FindAllUsers(c)
	if err != nil {
		return nil, err
	}

	userDTOs := make([]*entities.UserDTO, len(users))

	for i, user := range users {
		userDTOs[i] = &entities.UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Username: user.Username,
		}
	}

	return userDTOs, nil
}
