package usecases

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"gorm.io/gorm"
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
	var userIDAdmin uint
	userIDAdmin = 1
	var customer *entities.Customer
	_ = copier.Copy(&customer, &req)
	customer.Actor.CreatedByUserID = &userIDAdmin
	customer.Actor.UpdatedByUserID = &userIDAdmin
	err := u.CustomersRepo.Create(customer)
	if err != nil {
		logrus.Errorf("Create customer error: %v", err)
		return err
	}
	return nil
}

// Update update
func (u *customersUse) Update(req *entities.UpdateCustomerRequest) error {
	customer, err := u.CustomersRepo.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Get customer by id error: %v", err)
		return err
	}

	if customer == nil {
		logrus.Errorf("Customer not found")
		return gorm.ErrRecordNotFound
	}

	updateCustomer(req, customer)
	err = u.CustomersRepo.Update(customer)
	if err != nil {
		logrus.Errorf("Update customer error: %v", err)
		return err
	}
	return nil
}

func updateCustomer(req *entities.UpdateCustomerRequest, customer *entities.Customer) {
	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Surname != "" {
		customer.Surname = req.Surname
	}
	if req.Nickname != "" {
		customer.Nickname = req.Nickname
	}
	if req.TelephoneNumber != "" {
		customer.TelephoneNumber = req.TelephoneNumber
	}
	if req.PhoneNumber != "" {
		customer.PhoneNumber = req.PhoneNumber
	}
	if req.Detail != "" {
		customer.Detail = req.Detail
	}

	var userIDAdmin uint
	userIDAdmin = 1
	customer.Actor.UpdatedByUserID = &userIDAdmin
}

// GetAll get all
func (u *customersUse) GetAll(req *entities.GetAllCustomerRequest) ([]*entities.Customer, error) {
	customers, err := u.CustomersRepo.GetAll(req)
	if err != nil {
		logrus.Errorf("Get all customer error: %v", err)
		return nil, err
	}
	return customers, nil
}
