package dto

import (
	"time"

	"github.com/google/uuid"
)

type CourseGetResultRequestDto struct {
	CourseUID string `query:"uid"`
}

type CourseCreateResultRequestDto struct {
	CourseID               uuid.UUID `json:"uid"`
	CategoryName        string    `json:"category_name"`
	LearningOutcome     string    `json:"learning_outcome"`
	LearningExpectation string    `json:"learning_expectation"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type CourseUpdateResultRequestDto struct {
	CourseID               *uuid.UUID `json:"uid"`
	CategoryName        *string    `json:"category_name"`
	LearningOutcome     *string    `json:"learning_outcome"`
	LearningExpectation *string    `json:"learning_expectation"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type CourseResultResponseDto struct {
	ID                  *int       `json:"id"`
	CategoryName        *string    `json:"category_name"`
	LearningOutcome     *string    `json:"learning_outcome"`
	LearningExpectation *string    `json:"learning_expectation"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}
