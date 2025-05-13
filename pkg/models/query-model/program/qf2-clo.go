package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramCLOQueryEntity struct {
	ID              *uint           `gorm:"column:id"`
	PlanID          *uint           `gorm:"column:plan_id"`
	CLOGeneralData  *string         `gorm:"column:clo_general_data"`
	CLOSpecificData *string         `gorm:"column:clo_specific_data"`
	CreatedAt       *time.Time      `gorm:"column:created_at"`
	UpdatedAt       *time.Time      `gorm:"column:updated_at"`
	DeletedAt       *gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the Catagory6 model
func (q *ProgramCLOQueryEntity) TableName() string {
	return "program_clo"
}
