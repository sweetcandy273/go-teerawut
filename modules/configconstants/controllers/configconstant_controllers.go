package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	handlers "github.com/sweetcandy273/go-teerawut/pkg/handlers/response.go"
)

type configConstantController struct {
	ConfigConstantUse entities.ConfigConstantUsecase
}

// NewConfigConstantController is a constructor
func NewConfigConstantController(r fiber.Router, configConstantUse entities.ConfigConstantUsecase) {
	controllers := &configConstantController{
		ConfigConstantUse: configConstantUse,
	}
	r.Post("/", controllers.Create)
	r.Get("/", controllers.GetAll)
	// r.Patch("/:id", controllers.Update)
	// r.Delete("/:id", controllers.Delete)
}

// Create create
func (h *configConstantController) Create(c *fiber.Ctx) error {
	return handlers.ResponseSuccess(c, h.ConfigConstantUse.Create, &entities.CreateConfigConstantRequest{})
}

// GetAll get all
func (h *configConstantController) GetAll(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, h.ConfigConstantUse.GetAll, &entities.GetConfigConstantRequest{})
}
