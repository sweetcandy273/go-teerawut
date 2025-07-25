package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/render"
)

// WrapError wrap error
func WrapError() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return render.Error(c, err)
		}
		return nil
	}
}
