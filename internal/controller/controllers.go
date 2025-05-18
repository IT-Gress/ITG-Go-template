package controller

import "github.com/it-gress/itg-go-template/internal/repository"

// Controllers is a struct that holds all the controllers in the application.
type Controllers struct {
	UserController *UserController
}

// NewControllers initializes all controllers.
func NewControllers(repos *repository.Repositories) *Controllers {
	return &Controllers{
		UserController: NewUserController(repos.UserRepository),
	}
}
