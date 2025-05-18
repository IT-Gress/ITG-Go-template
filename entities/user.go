package entities

// UserDTO is a struct that represents a user data transfer object.
type UserDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	RoleID    string `json:"role_id"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
