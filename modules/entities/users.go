package entities

import "github.com/sweetcandy273/go-teerawut/pkg/handlers/context"

// UsersUsecase users usecase
type UsersUsecase interface {
	Register(c *context.Context, req *CreateUserRequest) (*User, error)
	Login(c *context.Context, req *LoginRequest) (*Token, error)
}

// UsersRepository users repository
type UsersRepository interface {
	Create(req *CreateUserRequest) (*User, error)
	FindByUsername(username string) (*User, error)
}

// User users register request
type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	DisplayName string `json:"display_name"`
}

// CreateUserRequest create request
type CreateUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
}

// LoginRequest login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
