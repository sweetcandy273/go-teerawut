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
	err := req.Validate()
	if err != nil {
		logrus.Errorf("Validate error: %v", err)
		return err
	}
	userIDAdmin := entities.UserIDAdmin
	customer := &entities.Customer{}
	_ = copier.CopyWithOption(&customer, req, copier.Option{IgnoreEmpty: true})
	customer.Actor = entities.Actor{
		CreatedByUserID: &userIDAdmin,
		UpdatedByUserID: &userIDAdmin,
	}
	err = u.CustomersRepo.Create(customer)
	if err != nil {
		logrus.Errorf("Create customer error: %v", err)
		return err
	}
	return nil
}

// Update update
func (u *customersUse) Update(req *entities.UpdateCustomerRequest) error {
	err := req.Validate()
	if err != nil {
		logrus.Errorf("Validate error: %v", err)
		return err
	}
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
	if req.PhoneNumber != "" {
		customer.PhoneNumber = req.PhoneNumber
	}
	if req.Detail != "" {
		customer.Detail = req.Detail
	}

	userIDAdmin := entities.UserIDAdmin
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

// Delete delete
func (u *customersUse) Delete(req *entities.GetOne) error {
	customer, err := u.CustomersRepo.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Get customer by id error: %v", err)
		return err
	}

	if customer == nil {
		logrus.Errorf("Customer not found")
		return gorm.ErrRecordNotFound
	}

	err = u.CustomersRepo.Delete(customer.ID)
	if err != nil {
		logrus.Errorf("Delete customer error: %v", err)
		return err
	}
	return nil
}

// GetByID get by id
func (u *customersUse) GetByID(req *entities.GetOne) (*entities.Customer, error) {
	customer, err := u.CustomersRepo.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Get customer by id %d error: %v", req, err)
		return nil, err
	}
	return customer, nil
}

// GetByDetailAndTelephoneNumber get by detail and telephone number
func (u *customersUse) GetByDetailAndTelephoneNumber(req *entities.GetByDetailAndTelephoneNumberRequest) (any, error) {
	customer, err := u.CustomersRepo.FindByDetailAndTelephoneNumber(req.Detail, req.TelephoneNumber)
	if err != nil {
		logrus.Errorf("Get customer by detail %s and telephone number %s error: %v", req.Detail, req.TelephoneNumber, err)
		return nil, err
	}
	return customer, nil
}
