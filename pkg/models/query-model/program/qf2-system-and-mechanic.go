package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramSystemAndMechanicQueryEntity struct {
	ID                      *uint       `gorm:"column:id"`
	CoursePolicies          *string     `gorm:"column:course_policies"`
	CourseStrategies        *string     `gorm:"column:course_strategies"`
	CourseRisk              *string     `gorm:"column:course_risk"`
	CourseStudentComment    *string     `gorm:"column:course_student_comment"`
	CourseExpectedAttribute *string     `gorm:"column:course_expected_attribute"`
	MainContentAndStructure *string     `gorm:"column:main_content_and_structure"`
	CourseImprovingPlan     *string     `gorm:"column:course_improving_plan"`
	CreatedAt               *time.Time  `gorm:"column:created_at"`
	UpdatedAt               *time.Time  `gorm:"column:updated_at"`
	DeletedAt               gorm.DeletedAt
}


func (q *ProgramSystemAndMechanicQueryEntity) TableName() string {
	return "program_system_and_mechanic"
}
