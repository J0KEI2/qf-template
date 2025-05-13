package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramPolicyAndStrategicQueryEntity struct {
	ID                *uint      `gorm:"column:id"`
	ProgramPhilosophy *string    `gorm:"column:program_philosophy"`
	ProgramObjective  *string    `gorm:"column:program_objective"`
	ProgramPolicy     *string    `gorm:"column:program_policy"`
	ProgramStrategic  *string    `gorm:"column:program_strategic"`
	ProgramRisk       *string    `gorm:"column:program_risk"`
	ProgramFeedback   *string    `gorm:"column:program_feedback"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at"`
	DeletedAt         *gorm.DeletedAt
}

// TableName sets the table name for the Catagory2 model
func (q *ProgramPolicyAndStrategicQueryEntity) TableName() string {
	return "program_policy_and_strategic"
}
