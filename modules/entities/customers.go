package entities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sweetcandy273/go-teerawut/pkg/utils"
	"gorm.io/gen"
)

// CustomersUsecase customers usecase
type CustomersUsecase interface {
	Create(req *CreateCustomerRequest) error
	Update(req *UpdateCustomerRequest) error
	GetAll(req *GetAllCustomerRequest) ([]*Customer, error)
	Delete(req *GetOne) error
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
}

// Customer customers register request
type Customer struct {
	Model
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Nickname        string `json:"nickname"`
	TelephoneNumber string `json:"telephone_number"`
	PhoneNumber     string `json:"phone_number"`
	Detail          string `json:"detail"`
	Actor
}

// CreateCustomerRequest create customer request
type CreateCustomerRequest struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Nickname        string `json:"nickname"`
	TelephoneNumber string `json:"telephone_number"`
	PhoneNumber     string `json:"phone_number"`
	Detail          string `json:"detail"`
}

// Validate validate
func (req CreateCustomerRequest) Validate() error {
	if req.TelephoneNumber == "" && req.PhoneNumber == "" {
		return fiber.NewError(fiber.StatusBadRequest, "telephone_number or phone_number is required")
	}
	if req.TelephoneNumber == "" && utils.IsValidTelephoneNumber(req.TelephoneNumber) {
		return fiber.NewError(fiber.StatusBadRequest, "telephone_number is invalid")
	}
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
	Surname         *string `json:"surname" query:"surname"`
	Nickname        *string `json:"nickname" query:"nickname"`
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
