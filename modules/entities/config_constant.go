package entities

import "github.com/sweetcandy273/go-teerawut/pkg/handlers/context"

// ConfigConstantRepository interface
type ConfigConstantRepository interface {
	Create(c *ConfigConstant) error
	GetAll(req *GetConfigConstantRequest) ([]*ConfigConstant, error)
}

// ConfigConstantUsecase interface
type ConfigConstantUsecase interface {
	Create(c *context.Context, req *CreateConfigConstantRequest) error
	GetAll(c *context.Context, req *GetConfigConstantRequest) ([]*ConfigConstant, error)
}

// ConfigConstant configuration constant entity
type ConfigConstant struct {
	Model
	Group       string `json:"group"`
	NameEn      string `json:"name_en"`
	NameTh      string `json:"name_th"`
	Option      string `json:"option"`
	Description string `json:"description"`
	Sort        uint   `json:"sort"`
}

// CreateConfigConstantRequest create configuration constant request
type CreateConfigConstantRequest struct {
	Group       string `json:"group" validate:"required"`
	NameTh      string `json:"name_th" validate:"required"`
	Option      string `json:"option"`
	Description string `json:"description"`
	Sort        uint   `json:"sort" validate:"required"`
}

// GetConfigConstantRequest request to get configuration constants
type GetConfigConstantRequest struct {
	Group string `json:"group" query:"group"`
}
