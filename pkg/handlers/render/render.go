package render

import (
	"github.com/gofiber/fiber/v2"
)

// JSON render json to client
func JSON(c *fiber.Ctx, response interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      response,
	})
}

// Byte render byte to client
func Byte(c *fiber.Ctx, bytes []byte) error {
	_, err := c.Status(fiber.StatusOK).
		Write(bytes)

	return err
}

// Error render error to client
func Error(c *fiber.Ctx, err error) error {
	return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
		"status":      fiber.ErrInternalServerError.Message,
		"status_code": fiber.ErrInternalServerError.Code,
		"message":     err.Error(),
		"result":      nil,
	})
}
