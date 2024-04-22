package entities

// UsersUsecase users usecase
type UsersUsecase interface {
	Register(req *CreateUserRequest) (*UserResponse, error)
}

// UsersRepository users repository
type UsersRepository interface {
	Register(req *CreateUserRequest) (*UserResponse, error)
}

// User users register request
type User struct {
	ID       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

// CreateUserRequest create request
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResponse users register response
type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

// LoginRequest login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
