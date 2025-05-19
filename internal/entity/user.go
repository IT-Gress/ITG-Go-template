package entity

// UserDTO is a struct that represents a user data transfer object.
type UserDTO struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	RoleID    int     `json:"role_id"`
	LastLogin *string `json:"last_login"`
	IsActive  bool    `json:"is_active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// CreateUserRequest is a struct that represents a request to create a new user.
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   int    `json:"role_id" binding:"required"`
}

// UpdateUserRequest is a struct that represents a request to update an existing user.
type UpdateUserRequest struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Username string  `json:"username" binding:"required"`
	Password *string `json:"password"`
	RoleID   int     `json:"role_id" binding:"required"`
	IsActive bool    `json:"is_active" binding:"required"`
}

// UserLoginRequest is a struct that represents a request to log in a user.
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
