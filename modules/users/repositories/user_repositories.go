package repositories

import (
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"github.com/sweetcandy273/go-teerawut/query"
	"gorm.io/gorm"
)

type usersRepo struct {
	DB *gorm.DB
}

// NewUsersRepository creates a new instance of UsersRepository
func NewUsersRepository(db *gorm.DB) entities.UsersRepository {
	return &usersRepo{
		DB: db,
	}
}

// Create create
func (r *usersRepo) Create(req *entities.CreateUserRequest) (*entities.UserResponse, error) {
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

// FindByUsername find user by username
func (r *usersRepo) FindByUsername(username string) (*entities.User, error) {
	query.SetDefault(r.DB)
	q := query.User
	user, err := q.Where(q.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
