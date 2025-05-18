package entities

// UserDTO is a struct that represents a user data transfer object.
type UserDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	RoleID    int    `json:"role_id"`
	LastLogin string `json:"last_login"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateUserRequest is a struct that represents a request to create a new user.
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   int    `json:"role_id" validate:"required"`
}

// UpdateUserRequest is a struct that represents a request to update an existing user.
type UpdateUserRequest struct {
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Username string  `json:"username" validate:"required"`
	Password *string `json:"password"`
	RoleID   int     `json:"role_id" validate:"required"`
	IsActive bool    `json:"is_active" validate:"required"`
}
