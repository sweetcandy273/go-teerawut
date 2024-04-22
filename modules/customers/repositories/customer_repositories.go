package repositories

import (
	"github.com/jinzhu/copier"
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
func (r *customersRepo) Create(req *entities.CreateCustomerRequest) error {
	var userIDAdmin uint
	userIDAdmin = 1
	var customer entities.Customer
	_ = copier.Copy(&customer, &req)
	customer.Actor.CreatedByUserID = &userIDAdmin
	customer.Actor.UpdatedByUserID = &userIDAdmin

	err := r.DB.Create(&customer).Error
	if err != nil {
		logrus.Errorf("Create customer error: %v", err)
		return err
	}

	return nil
}
