package dto

import (
	"time"

	"gorm.io/gorm"
)

type ProgramSystemAndMechanicDto struct {
	ID                      *uint           `json:"id"`
	CoursePolicies          *string         `json:"course_policies"`
	CourseStrategies        *string         `json:"course_strategies"`
	CourseRisk              *string         `json:"course_risk"`
	CourseStudentComment    *string         `json:"course_student_comment"`
	CourseExpectedAttribute *string         `json:"course_expected_attribute"`
	MainContentAndStructure *string         `json:"main_content_and_structure"`
	CourseImprovingPlan     *string         `json:"course_improving_plan"`
	CreatedAt               *time.Time      `json:"created_at"`
	UpdatedAt               *time.Time      `json:"updated_at"`
	DeletedAt               *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type CreateOrUpdateSystemAndMechanicDto struct {
	ID                      *uint           `json:"id"`
	CoursePolicies          *string         `json:"course_policies"`
	CourseStrategies        *string         `json:"course_strategies"`
	CourseRisk              *string         `json:"course_risk"`
	CourseStudentComment    *string         `json:"course_student_comment"`
	CourseExpectedAttribute *string         `json:"course_expected_attribute"`
	MainContentAndStructure *string         `json:"main_content_and_structure"`
	CourseImprovingPlan     *string         `json:"course_improving_plan"`
	CreatedAt               *time.Time      `json:"created_at"`
	UpdatedAt               *time.Time      `json:"updated_at"`
	DeletedAt               *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
