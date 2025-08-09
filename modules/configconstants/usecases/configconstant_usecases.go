package usecases

import (
	"github.com/sweetcandy273/go-teerawut/pkg/handlers/context"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"github.com/sweetcandy273/go-teerawut/modules/entities"
)

type configConstantUse struct {
	ConfigConstantRepo entities.ConfigConstantRepository
}

// Constructor
func NewConfigConstantUsecase(configConstantRepo entities.ConfigConstantRepository) entities.ConfigConstantUsecase {
	return &configConstantUse{
		ConfigConstantRepo: configConstantRepo,
	}
}

// Create create
func (u *configConstantUse) Create(c *context.Context, req *entities.CreateConfigConstantRequest) error {
	configConstant := &entities.ConfigConstant{}
	_ = copier.CopyWithOption(&configConstant, req, copier.Option{IgnoreEmpty: true})
	err := u.ConfigConstantRepo.Create(configConstant)
	if err != nil {
		logrus.Errorf("Create config constant error: %v", err)
		return err
	}

	return nil
}
