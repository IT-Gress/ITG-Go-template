package repository

import "github.com/jmoiron/sqlx"

// Repositories is a struct that holds all the repositories in the application.
type Repositories struct {
	UserRepository *UserRepository
}

// NewRepositories initializes all repositories.
func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
