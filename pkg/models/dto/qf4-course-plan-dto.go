package dto

import (
	"time"

	"github.com/google/uuid"
)

type QF4GetCoursePlanRequestDto struct {
	QF4UID string `query:"uid"`
}

type QF4CreateCoursePlanRequestDto struct {
	ID               *int       `json:"id"`
	Qf4MainUid       *uuid.UUID `json:"qf4_main_uuid"`
	Week             *string    `json:"week"`
	Title            *string    `json:"title"`
	PlanDescription  *string    `json:"other"`
	TheoryHour       *int       `json:"theory_hour"`
	OperationHour    *int       `json:"operation_hour"`
	SelfLearningHour *int       `json:"self_learning_hour"`
	LearningOutcome  *string    `json:"learning_outcome"`
	TeachingMedia    *string    `json:"teaching_outcome"`
	LeaningActivity  *string    `json:"learning_activity"`
	EvaluationMethod *string    `json:"evaluation_method"`
	Lecturer         *string    `json:"lecturer"`
	AssessmentScore  *int       `json:"assessment_score"`
	AssessmentNote   *string    `json:"assessment_note"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type QF4UpdateCoursePlanRequestDto struct {
	ID               *int       `json:"id"`
	Qf4MainUid       *uuid.UUID `json:"qf4_main_uuid"`
	Week             *string    `json:"week"`
	Title            *string    `json:"title"`
	PlanDescription  *string    `json:"other"`
	TheoryHour       *int       `json:"theory_hour"`
	OperationHour    *int       `json:"operation_hour"`
	SelfLearningHour *int       `json:"self_learning_hour"`
	LearningOutcome  *string    `json:"learning_outcome"`
	TeachingMedia    *string    `json:"teaching_outcome"`
	LeaningActivity  *string    `json:"learning_activity"`
	EvaluationMethod *string    `json:"evaluation_method"`
	Lecturer         *string    `json:"lecturer"`
	AssessmentScore  *int       `json:"assessment_score"`
	AssessmentNote   *string    `json:"assessment_note"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type QF4CoursePlanResponseDto struct {
	ID               *int       `json:"id" sql:"AUTO_INCREMENT"`
	Qf4MainUid       *uuid.UUID `json:"qf4_main_uuid"`
	Week             *string    `json:"week"`
	Title            *string    `json:"title"`
	PlanDescription  *string    `json:"other"`
	TheoryHour       *int       `json:"theory_hour"`
	OperationHour    *int       `json:"operation_hour"`
	SelfLearningHour *int       `json:"self_learning_hour"`
	LearningOutcome  *string    `json:"learning_outcome"`
	TeachingMedia    *string    `json:"teaching_outcome"`
	LeaningActivity  *string    `json:"learning_activity"`
	EvaluationMethod *string    `json:"evaluation_method"`
	Lecturer         *string    `json:"lecturer"`
	AssessmentScore  *int       `json:"assessment_score"`
	AssessmentNote   *string    `json:"assessment_note"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}
