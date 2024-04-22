package usecases

import (
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
)

type customersUse struct {
	CustomersRepo entities.CustomersRepository
}

// Constructor
func NewCustomersUsecase(customersRepo entities.CustomersRepository) entities.CustomersUsecase {
	return &customersUse{
		CustomersRepo: customersRepo,
	}
}

// Create create
func (u *customersUse) Create(req *entities.CreateCustomerRequest) error {
	err := u.CustomersRepo.Create(req)
	if err != nil {
		logrus.Errorf("Create customer error: %v", err)
		return err
	}
	return nil
}
