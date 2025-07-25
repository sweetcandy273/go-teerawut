package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	handlers "github.com/sweetcandy273/go-teerawut/pkg/handlers/response.go"
)

type usersController struct {
	UsersUse entities.UsersUsecase
}

// NewUsersController is a constructor
func NewUsersController(r fiber.Router, usersUse entities.UsersUsecase) {
	controllers := &usersController{
		UsersUse: usersUse,
	}
	r.Post("/", controllers.Register)
	r.Post("/login", controllers.Login)
}

// Register register
func (h *usersController) Register(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, h.UsersUse.Register, &entities.CreateUserRequest{})
}

// Login login
func (h *usersController) Login(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, h.UsersUse.Login, &entities.LoginRequest{})
}
