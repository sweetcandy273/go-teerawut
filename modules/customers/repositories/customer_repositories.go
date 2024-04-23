package repositories

import (
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"gorm.io/gorm"
)

type customersRepo struct {
	DB *gorm.DB
}

// NewCustomersRepository new customers repository
func NewCustomersRepository(db *gorm.DB) entities.CustomersRepository {
	return &customersRepo{
		DB: db,
	}
}

// Create create
func (r *customersRepo) Create(c *entities.Customer) error {
	err := r.DB.Create(&c).Error
	if err != nil {
		logrus.Errorf("Create customer error: %v", err)
		return err
	}

	return nil
}

// GetByID get by id
func (r *customersRepo) GetByID(id uint) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.DB.Where("id = ?", id).First(&customer).Error
	if err != nil {
		logrus.Errorf("Get customer by id error: %v", err)
		return nil, err
	}

	return &customer, nil
}

// Update update
func (r *customersRepo) Update(c *entities.Customer) error {
	err := r.DB.Save(&c).Error
	if err != nil {
		logrus.Errorf("Update customer error: %v", err)
		return err
	}

	return nil
}
