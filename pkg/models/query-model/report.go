package query

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReportQueryEntity struct {
	ID          *uint           `gorm:"column:id;type:bigint;primaryKey"`
	QFMainID    *uuid.UUID      `gorm:"column:qf_main_id;type:uuid"`
	Name        *string         `gorm:"column:name;type:varchar(255)"`
	Description *string         `gorm:"column:description;type:varchar(255)"`
	CreatedAt   *time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt   *gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

func (ReportQueryEntity) TableName() string {
	return "reports"
}
