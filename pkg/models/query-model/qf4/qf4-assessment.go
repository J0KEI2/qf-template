package query

import "time"

type QF4Assessment struct {
	ID                 *int
	CategoryName       *string
	LearningAssessment *int
	Grade              *string
	GroupBased         *string
	Other              *string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

type QF4AssessmentQueryEntity struct {
	ID                 *int
	CategoryName       *string
	LearningAssessment *int
	Grade              *string
	GroupBased         *string
	Other              *string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}