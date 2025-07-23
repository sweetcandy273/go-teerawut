package usecases

import (
	"errors"
	"fmt"

	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

// NewUsersUsecase creates a new instance of UsersUsecase
func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

// Register register
func (u *usersUse) Register(req *entities.CreateUserRequest) (*entities.UserResponse, error) {
	// Check if user already exists
	_, err := u.UsersRepo.FindByUsername(req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err.Error())
		return nil, err
	} else if err == nil {
		return nil, fmt.Errorf("user with username %s already exists", req.Username)
	}

	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.UsersRepo.Create(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
