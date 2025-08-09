package repository

import (
	"github.com/sweetcandy273/go-teerawut/modules/entities"
	"github.com/sweetcandy273/go-teerawut/query"
	"gorm.io/gen"
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

// GetAll get all
func (r *configConstantRepo) GetAll(req *entities.GetConfigConstantRequest) ([]*entities.ConfigConstant, error) {
	query.SetDefault(r.DB)
	q := query.ConfigConstant
	var conditions []gen.Condition
	if req.Group != "" {
		conditions = append(conditions, q.Group_.Eq(req.Group))
	}
	configConstants, err := q.Where(conditions...).Order(q.Sort).Find()
	if err != nil {
		return nil, err
	}
	return configConstants, nil
}
