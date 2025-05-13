package query

import (
	"time"

	"gorm.io/gorm"
)

type ReferenceOptionQueryEntity struct {
	ID        *int            `json:"id"`
	Name      *string         `json:"name"`
	UpdatedAt *time.Time      `json:"updated_at"`
	CreatedAt *time.Time      `json:"created_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

func (ReferenceOptionQueryEntity) TableName() string {
	return "common_reference_options"
}
