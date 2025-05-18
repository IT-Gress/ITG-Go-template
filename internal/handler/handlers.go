package handler

import "github.com/it-gress/itg-go-template/internal/controller"

// Handlers is a struct that holds all the handlers in the application.
type Handlers struct {
	UserHandler *UserHandler
}

// NewHandlers initializes all Handlers.
func NewHandlers(controllers *controller.Controllers) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(controllers.UserController),
	}
}
