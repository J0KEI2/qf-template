package dto

import (
	"time"

	"gorm.io/gorm"
)

type ProgramPlanAndEvaluateRequestDto struct {
	ID                                *uint       `json:"id"`
	ProgramSubPlanID                  *uint       `json:"program_sub_plan_id"`
	StudentCharacteristics            interface{} `json:"student_characteristics"`
	ReceiveStudentPlan                interface{} `json:"receive_student_plan"`
	ProgramIncome                     interface{} `json:"program_income"`
	ProgramOutcome                    interface{} `json:"program_outcome"`
	AcademicEvaluation                *string     `json:"academic_evaluation"`
	GraduationCriteria                *string     `json:"graduation_criteria"`
	ProgramUniversityTransferStandard *string     `json:"program_university_transfer_standard"`
	ProgramPreparation                *string     `json:"program_preparation"`
}

type ProgramPlanAndEvaluateResponseDto struct {
	ID                                *uint           `json:"id"`
	ProgramSubPlanID                  *uint           `json:"program_sub_plan_id"`
	ProgramDegreeTypeID               *int            `json:"program_degree_type_id"`
	ProgramDegreeType                 *string         `json:"program_degree_type"`
	ProgramYearID                     *int            `json:"program_year_id"`
	ProgramYear                       *uint           `json:"program_year"`
	StudentCharacteristics            interface{}     `json:"student_characteristics"`
	ReceiveStudentPlan                interface{}     `json:"receive_student_plan"`
	ProgramIncome                     interface{}     `json:"program_income"`
	ProgramOutcome                    interface{}     `json:"program_outcome"`
	AcademicEvaluation                *string         `json:"academic_evaluation"`
	GraduationCriteria                *string         `json:"graduation_criteria"`
	ProgramUniversityTransferStandard *string         `json:"program_university_transfer_standard"`
	ProgramPreparation                *string         `json:"program_preparation"`
	CreatedAt                         *time.Time      `json:"created_at"`
	UpdatedAt                         *time.Time      `json:"updated_at"`
	DeletedAt                         *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
