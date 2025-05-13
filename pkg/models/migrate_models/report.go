package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	ID          uint           `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	QFMainID    uuid.UUID      `gorm:"column:qf_main_id;type:uuid" json:"qf_main_id"`
	Name        string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Description string         `gorm:"column:description;type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at,omitempty"`
}
