package repository

import (
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"github.com/sweetcandy273/go-teerawut/query"
	"gorm.io/gorm"
)

type configConstantRepo struct {
	DB *gorm.DB
}

// NewConfigConstantRepository new config constant repository
func NewConfigConstantRepository(db *gorm.DB) entities.ConfigConstantRepository {
	return &configConstantRepo{
		DB: db,
	}
}

// Create create
func (r *configConstantRepo) Create(c *entities.ConfigConstant) error {
	query.SetDefault(r.DB)
	q := query.ConfigConstant
	return q.Create(c)
}
