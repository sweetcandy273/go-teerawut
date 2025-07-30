package entities

import (
	"time"

	"gorm.io/gorm"
)

// Actor actor
type Actor struct {
	CreatedByUserID uint  `json:"create_by_user_id" db:"create_by_user_id"`
	UpdatedByUserID *uint `json:"update_by_user_id" db:"update_by_user_id"`
	DeletedByUserID *uint `json:"delete_by_user_id" db:"delete_by_user_id"`
}

// Model base model
type Model struct {
	ID        uint           `gorm:"primaryKey;autoIncrement:true;unique" json:"id" column:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// GetOne get one
type GetOne struct {
	ID uint `json:"-" path:"id" form:"id" query:"id" validate:"required"`
}
