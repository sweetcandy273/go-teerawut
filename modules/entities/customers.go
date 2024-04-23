package entities

// CustomersUsecase customers usecase
type CustomersUsecase interface {
	Create(req *CreateCustomerRequest) error
	Update(req *UpdateCustomerRequest) error
}

// CustomersRepository customers repository
type CustomersRepository interface {
	Create(c *Customer) error
	Update(c *Customer) error
	GetByID(id uint) (*Customer, error)
}

// Customer customers register request
type Customer struct {
	Model
	Name            string `json:"name" db:"name"`
	Surname         string `json:"surname" db:"surname"`
	Nickname        string `json:"nickname" db:"nickname"`
	TelephoneNumber string `json:"telephone_number" db:"telephone_number"`
	PhoneNumber     string `json:"phone_number" db:"phone_number"`
	Detail          string `json:"detail" db:"detail"`
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

// UpdateCustomerRequest update customer request
type UpdateCustomerRequest struct {
	ID uint `json:"-" path:"id" form:"id" query:"id" validate:"required"`
	CreateCustomerRequest
}