package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramPlanAndEvaluateQueryEntity struct {
	ID                                *uint                      `gorm:"column:id"`
	ProgramSubPlan                    *ProgramSubPlanQueryEntity `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID                  *uint                      `gorm:"column:program_sub_plan_id;type:uint"`
	StudentCharacteristic             *string                    `gorm:"column:student_characteristic"`
	ReceiveStudentPlan                *string                    `gorm:"column:receive_student_plan"`
	ProgramIncome                     *string                    `gorm:"column:program_income"`
	ProgramOutcome                    *string                    `gorm:"column:program_outcome"`
	AcademicEvaluation                *string                    `gorm:"column:academic_evaluation"`
	GraduationCriteria                *string                    `gorm:"column:graduation_criteria"`
	ProgramUniversityTransferStandard *string                    `gorm:"column:program_university_transfer_standard"`
	ProgramPreparation                *string                    `gorm:"column:program_preparation"`
	CreatedAt                         *time.Time                 `gorm:"column:created_at"`
	UpdatedAt                         *time.Time                 `gorm:"column:updated_at"`
	DeletedAt                         *gorm.DeletedAt
}

// TableName specifies the table name for the category_8 struct
func (q *ProgramPlanAndEvaluateQueryEntity) TableName() string {
	return "program_plan_and_evaluate"
}
