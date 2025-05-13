package dto

import (
	"time"

	"github.com/google/uuid"
)

type QF4GetAssessmentRequestDto struct {
	QF4UID string `query:"uid"`
}

type QF4CreateAssessmentRequestDto struct {
	QF4ID              uuid.UUID `json:"uid"`
	CategoryName       string    `json:"category_name"`
	LearningAssessment int       `json:"learning_assessment"`
	Grade              string    `json:"grade"`
	GroupBased         string    `json:"group_based"`
	Other              string    `json:"other"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type QF4UpdateAssessmentRequestDto struct {
	QF4ID              uuid.UUID `json:"uid"`
	CategoryName       string    `json:"category_name"`
	LearningAssessment int       `json:"learning_assessment"`
	Grade              string    `json:"grade"`
	GroupBased         string    `json:"group_based"`
	Other              string    `json:"other"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type QF4AssessmentResponseDto struct {
	ID                 *int       `json:"id"`
	CategoryName       *string    `json:"category_name"`
	LearningAssessment *int       `json:"learning_assessment"`
	Grade              *string    `json:"grade"`
	GroupBased         *string    `json:"group_based"`
	Other              *string    `json:"other"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}
