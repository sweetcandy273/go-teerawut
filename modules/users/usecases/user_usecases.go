package usecases

import (
	"fmt"

	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"golang.org/x/crypto/bcrypt"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

// Constructor
func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

// Register register
func (u *usersUse) Register(req *entities.CreateUserRequest) (*entities.UserResponse, error) {
	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.UsersRepo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
