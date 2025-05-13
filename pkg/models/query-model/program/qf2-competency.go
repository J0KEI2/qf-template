package query

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramCompetencyQueryEntity struct {
	ID                 *uint                   `gorm:"column:id"`
	ProgramMainID      *uuid.UUID              `gorm:"column:program_main_id"`
	ProgramMain        *ProgramMainQueryEntity `gorm:"foreignKey:ProgramMainID;references:ID"`
	Order              *int                    `gorm:"column:order"`
	SpecificCompetency *string                 `gorm:"column:specific_competency"`
	GenericCompetency  *string                 `gorm:"column:generic_competency"`
	CreatedAt          *time.Time              `gorm:"column:created_at"`
	UpdatedAt          *time.Time              `gorm:"column:updated_at"`
	DeletedAt          *gorm.DeletedAt         `gorm:"column:deleted_at"`
}

// TableName sets the table name for the Catagory6 model
func (q *ProgramCompetencyQueryEntity) TableName() string {
	return "program_competency"
}
