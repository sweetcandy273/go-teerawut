package repositories

import (
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"gorm.io/gorm"
)

type usersRepo struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) entities.UsersRepository {
	return &usersRepo{
		DB: db,
	}
}

// Register register
func (r *usersRepo) Register(req *entities.CreateUserRequest) (*entities.UserResponse, error) {
	user := entities.User{
		Username: req.Username,
		Password: req.Password,
	}
	err := r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &entities.UserResponse{
		Model:    user.Model,
		Username: user.Username,
	}, nil
}
