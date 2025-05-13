package dto

import (
	"time"

	"github.com/google/uuid"
)

type QF4GetResultRequestDto struct {
	QF4UID string `query:"uid"`
}

type QF4CreateResultRequestDto struct {
	QF4ID               uuid.UUID `json:"uid"`
	CategoryName        string    `json:"category_name"`
	LearningOutcome     string    `json:"learning_outcome"`
	LearningExpectation string    `json:"learning_expectation"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type QF4UpdateResultRequestDto struct {
	QF4ID               *uuid.UUID `json:"uid"`
	CategoryName        *string    `json:"category_name"`
	LearningOutcome     *string    `json:"learning_outcome"`
	LearningExpectation *string    `json:"learning_expectation"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type QF4ResultResponseDto struct {
	ID                  *int       `json:"id"`
	CategoryName        *string    `json:"category_name"`
	LearningOutcome     *string    `json:"learning_outcome"`
	LearningExpectation *string    `json:"learning_expectation"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}
