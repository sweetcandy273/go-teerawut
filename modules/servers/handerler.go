package servers

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	_usersHttp "github.com/sweetcandy273/go-teerawut/modules/users/controllers"
	_usersRepository "github.com/sweetcandy273/go-teerawut/modules/users/repositories"
	_usersUsecase "github.com/sweetcandy273/go-teerawut/modules/users/usecases"
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/context"
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/middlewares"

	_customersHttp "github.com/sweetcandy273/go-teerawut/modules/customers/controllers"
	_customersRepository "github.com/sweetcandy273/go-teerawut/modules/customers/repositories"
	_customersUsecase "github.com/sweetcandy273/go-teerawut/modules/customers/usecases"

	_configConstantHttp "github.com/sweetcandy273/go-teerawut/modules/configconstants/controllers"
	_configConstantRepository "github.com/sweetcandy273/go-teerawut/modules/configconstants/repositories"
	_configConstantUsecase "github.com/sweetcandy273/go-teerawut/modules/configconstants/usecases"

	"github.com/gofiber/fiber/v2"
)

// MapHandlers map handlers
func (s *Server) MapHandlers() error {
	s.App.Use(
		middlewares.WrapError(), // Wrap error middleware
		cors.New(
			cors.Config{
				AllowOrigins: "*", // Allow all origins
				AllowMethods: "GET, POST, PUT, DELETE, OPTIONS, PATCH",
				AllowHeaders: "Origin, Content-Type, Accept, Authorization",
			},
		),
	)

	authJWT := jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")), // Secret key for signing JWT
		ContextKey: context.UserKey,                 // จะเก็บ token ใน c.Locals("user")
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	})

	// Group a version
	v1 := s.App.Group("/v1")

	//* Health check group
	healthCheckGroup := v1.Group("/health-check")
	healthCheckGroup.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":      "OK",
			"status_code": fiber.StatusOK,
			"message":     "server is running",
			"result":      nil,
		})
	})

	//* Config constants group
	configConstantGroup := v1.Group("/config_constants", authJWT)
	configConstantRepository := _configConstantRepository.NewConfigConstantRepository(s.DB)
	configConstantUsecase := _configConstantUsecase.NewConfigConstantUsecase(configConstantRepository)
	_configConstantHttp.NewConfigConstantController(configConstantGroup, configConstantUsecase)

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.DB)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	//* Customers group
	customersGroup := v1.Group("/customers", authJWT)
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
