package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID           int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	QFType       string         `gorm:"column:qf_type;type:varchar;size:255"`
	QFUID        uuid.UUID      `gorm:"column:qf_uid;type:uuid;"`
	CategoryType string         `gorm:"column:category_type;type:varchar;size:255"`
	Attribute    string         `gorm:"column:attribute;type:varchar;size:255"`
	Commentator  string         `gorm:"column:commentator;type:varchar;size:255"`
	Comments     string         `gorm:"column:comments;type:varchar;size:255"`
	Resolve      bool           `gorm:"column:resolve;type:boolean"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index;column:deleted_at"`
	UpdatedBy    uuid.UUID      `gorm:"column:updated_by;type:uuid;not null"`
}
