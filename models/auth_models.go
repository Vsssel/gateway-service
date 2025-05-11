package models

// LoginRequest defines the structure for user login
type LoginRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"securepassword"`
}

// LoginRequest defines the structure for user register
type RegisterRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"securepassword"`
	Email    string `json:"email"    example:"securepassword`
	RoleID   int    `json:"role_id " example:1`
}

// LoginResponse defines the structure for login response
type AuthResponse struct {
	AccessToken    string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// ErrorResponse defines the structure for error responses
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}