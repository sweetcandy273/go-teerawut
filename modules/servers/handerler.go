package servers

import (
	_usersHttp "github.com/sweetcandy273/go-teerawut/modules/users/controllers"
	_usersRepository "github.com/sweetcandy273/go-teerawut/modules/users/repositories"
	_usersUsecase "github.com/sweetcandy273/go-teerawut/modules/users/usecases"

	_customersHttp "github.com/sweetcandy273/go-teerawut/modules/customers/controllers"
	_customersRepository "github.com/sweetcandy273/go-teerawut/modules/customers/repositories"
	_customersUsecase "github.com/sweetcandy273/go-teerawut/modules/customers/usecases"

	"github.com/gofiber/fiber/v2"
)

// MapHandlers map handlers
func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.DB)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	//* Customers group
	customersGroup := v1.Group("/customers")
	customersRepository := _customersRepository.NewCustomersRepository(s.DB)
	customersUsecase := _customersUsecase.NewCustomersUsecase(customersRepository)
	_customersHttp.NewCustomersController(customersGroup, customersUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
