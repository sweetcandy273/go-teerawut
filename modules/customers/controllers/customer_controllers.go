package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	handlers "github.com/sweetcandy273/go-teerawut/pkg/handlers/response.go"
)

type customersController struct {
	CustomersUse entities.CustomersUsecase
}

// NewUsersController is a constructor
func NewCustomersController(r fiber.Router, customerUse entities.CustomersUsecase) {
	controllers := &customersController{
		CustomersUse: customerUse,
	}
	r.Post("/", controllers.Create)
	r.Patch("/:id", controllers.Update)
	r.Get("/", controllers.GetAll)
	r.Delete("/:id", controllers.Delete)
	r.Get("/detail_and_telephone_number", controllers.GetByDetailAndTelephoneNumber)
	r.Get("/:id", controllers.GetByID)
}

// Create create
func (h *customersController) Create(c *fiber.Ctx) error {
	return handlers.ResponseSuccess(c, h.CustomersUse.Create, &entities.CreateCustomerRequest{})
}

// Update update
func (h *customersController) Update(c *fiber.Ctx) error {
	return handlers.ResponseSuccess(c, h.CustomersUse.Update, &entities.UpdateCustomerRequest{})
}

// GetAll get all
func (h *customersController) GetAll(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, h.CustomersUse.GetAll, &entities.GetAllCustomerRequest{})
}

// Delete delete
func (h *customersController) Delete(c *fiber.Ctx) error {
	return handlers.ResponseSuccess(c, h.CustomersUse.Delete, &entities.GetOne{})
}

// GetByID get by id
func (h *customersController) GetByID(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, h.CustomersUse.GetByID, &entities.GetOne{})
}

// GetByDetailAndTelephoneNumber get by detail and telephone number
func (h *customersController) GetByDetailAndTelephoneNumber(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, h.CustomersUse.GetByDetailAndTelephoneNumber, &entities.GetByDetailAndTelephoneNumberRequest{})
}
