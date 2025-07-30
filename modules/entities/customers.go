package entities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/context"
	"github.com/sweetcandy273/go-teerawut/pkg/utils"
	"gorm.io/gen"
)

// CustomersUsecase customers usecase
type CustomersUsecase interface {
	Create(c *context.Context, req *CreateCustomerRequest) error
	Update(c *context.Context, req *UpdateCustomerRequest) error
	GetAll(req *GetAllCustomerRequest) ([]*Customer, error)
	Delete(c *context.Context, req *GetOne) error
	GetByID(req *GetOne) (*Customer, error)
	GetByDetailAndTelephoneNumber(req *GetByDetailAndTelephoneNumberRequest) (any, error)
}

// CustomersRepository customers repository
type CustomersRepository interface {
	Create(c *Customer) error
	Update(c *Customer) error
	GetByID(id uint) (*Customer, error)
	GetAll(req *GetAllCustomerRequest) ([]*Customer, error)
	Delete(id uint) error
	FindByDetailAndTelephoneNumber(detail, telephoneNumber string) (gen.T, error)
	CreateAddress(addresses []*CustomerAddress) error
	UpdateAddress(addresses []*CustomerAddress) error
	DeleteAddress(ids []uint) error
}

// Customer customers register request
type Customer struct {
	Model
	Name        string             `json:"name"`
	PhoneNumber string             `json:"phone_number"`
	Detail      string             `json:"detail"`
	Addresses   []*CustomerAddress `json:"addresses" gorm:"foreignKey:CustomerID;references:ID"`
	Actor
}

// CreateCustomerRequest create customer request
type CreateCustomerRequest struct {
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Detail      string    `json:"detail"`
	Addresses   []address `json:"addresses"`
}

type address struct {
	ID      uint   `json:"id"`
	Address string `json:"address"`
	Village string `json:"village"`
	Detail  string `json:"detail"`
}

// Validate validate
func (req CreateCustomerRequest) Validate() error {
	if req.PhoneNumber == "" && utils.IsValidPhoneNumber(req.PhoneNumber) {
		return fiber.NewError(fiber.StatusBadRequest, "phone_number is invalid")
	}
	return nil
}

// UpdateCustomerRequest update customer request
type UpdateCustomerRequest struct {
	ID uint `json:"-" path:"id" form:"id" query:"id" validate:"required"`
	CreateCustomerRequest
}

// GetAllCustomerRequest get all customer request
type GetAllCustomerRequest struct {
	ID              *uint   `json:"id" query:"id"`
	Name            *string `json:"name" query:"name"`
	TelephoneNumber *string `json:"telephone_number" query:"telephone_number"`
	PhoneNumber     *string `json:"phone_number" query:"phone_number"`
	Detail          *string `json:"detail" query:"detail"`
	Query           *string `json:"query" query:"query"`
}

// GetByDetailAndTelephoneNumberRequest get by detail and telephone number request
type GetByDetailAndTelephoneNumberRequest struct {
	Detail          string `json:"detail" query:"detail" validate:"required"`
	TelephoneNumber string `json:"telephone_number" query:"telephone_number" validate:"required"`
}
