package query

import (
	"time"

	"github.com/google/uuid"
)

type QF4CoursePlan struct {
	ID               int       `gorm:"column:id"`
	Qf4MainUid       uuid.UUID `gorm:"column:qf4_main_uuid"`
	Week             string    `gorm:"column:week"`
	Title            string    `gorm:"column:title"`
	PlanDescription  string    `gorm:"column:other"`
	TheoryHour       int       `gorm:"column:theory_hour"`
	OperationHour    int       `gorm:"column:operation_hour"`
	SelfLearningHour int       `gorm:"column:self_learning_hour"`
	LearningOutcome  string    `gorm:"column:learning_outcome"`
	TeachingMedia    string    `gorm:"column:teaching_outcome"`
	LeaningActivity  string    `gorm:"column:learning_activity"`
	EvaluationMethod string    `gorm:"column:evaluation_method"`
	Lecturer         string    `gorm:"column:lecturer"`
	AssessmentScore  int       `gorm:"column:assessment_score"`
	AssessmentNote   string    `gorm:"column:assessment_note"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

type QF4CoursePlanQueryEntity struct {
	ID               *int       `gorm:"column:id"`
	Qf4MainUid       *uuid.UUID `gorm:"column:qf4_main_uuid"`
	Week             *string    `gorm:"column:week"`
	Title            *string    `gorm:"column:title"`
	PlanDescription  *string    `gorm:"column:other"`
	TheoryHour       *int       `gorm:"column:theory_hour"`
	OperationHour    *int       `gorm:"column:operation_hour"`
	SelfLearningHour *int       `gorm:"column:self_learning_hour"`
	LearningOutcome  *string    `gorm:"column:learning_outcome"`
	TeachingMedia    *string    `gorm:"column:teaching_outcome"`
	LeaningActivity  *string    `gorm:"column:learning_activity"`
	EvaluationMethod *string    `gorm:"column:evaluation_method"`
	Lecturer         *string    `gorm:"column:lecturer"`
	AssessmentScore  *int       `gorm:"column:assessment_score"`
	AssessmentNote   *string    `gorm:"column:assessment_note"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at"`
}

func (s *QF4CoursePlanQueryEntity) TableName() string {
	return "qf4_course_plans"
}
