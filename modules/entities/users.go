package entities

const (
	// UserIDAdmin user id admin
	UserIDAdmin uint = 1
)

// UsersUsecase users usecase
type UsersUsecase interface {
	Register(req *CreateUserRequest) (*User, error)
	Login(req *LoginRequest) (*Token, error)
}

// UsersRepository users repository
type UsersRepository interface {
	Create(req *CreateUserRequest) (*User, error)
	FindByUsername(username string) (*User, error)
}

// User users register request
type User struct {
	Model
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
