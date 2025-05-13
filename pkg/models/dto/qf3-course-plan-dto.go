package dto

import (
	"time"

	"github.com/google/uuid"
)

type CourseGetCoursePlanRequestDto struct {
	CourseUID string `query:"uid"`
}

type CourseCreateCoursePlanRequestDto struct {
	ID               *int       `json:"id"`
	Qf3MainUid       *uuid.UUID `json:"course_main_uuid"`
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

type CourseUpdateCoursePlanRequestDto struct {
	ID               *int       `json:"id"`
	Qf3MainUid       *uuid.UUID `json:"course_main_uuid"`
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

type CoursePlanResponseDto struct {
	ID               *int       `json:"id" sql:"AUTO_INCREMENT"`
	Qf3MainUid       *uuid.UUID `json:"course_main_uuid"`
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
