package dto

import (
	"time"

	"github.com/google/uuid"
)

type CourseGetAssessmentRequestDto struct {
	CourseUID string `query:"uid"`
}

type CourseCreateAssessmentRequestDto struct {
	CourseID              uuid.UUID `json:"uid"`
	CategoryName       string    `json:"category_name"`
	LearningAssessment int       `json:"learning_assessment"`
	Grade              string    `json:"grade"`
	GroupBased         string    `json:"group_based"`
	Other              string    `json:"other"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CourseUpdateAssessmentRequestDto struct {
	CourseID              uuid.UUID `json:"uid"`
	CategoryName       string    `json:"category_name"`
	LearningAssessment int       `json:"learning_assessment"`
	Grade              string    `json:"grade"`
	GroupBased         string    `json:"group_based"`
	Other              string    `json:"other"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CourseAssessmentResponseDto struct {
	ID                 *int       `json:"id"`
	CategoryName       *string    `json:"category_name"`
	LearningAssessment *int       `json:"learning_assessment"`
	Grade              *string    `json:"grade"`
	GroupBased         *string    `json:"group_based"`
	Other              *string    `json:"other"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}
