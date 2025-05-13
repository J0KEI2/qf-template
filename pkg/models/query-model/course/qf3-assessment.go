package query

import "time"

type CourseAssessment struct {
	ID                 *int
	CategoryName       *string
	LearningAssessment *int
	Grade              *string
	GroupBased         *string
	Other              *string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

type CourseAssessmentQueryEntity struct {
	ID                 *int
	CategoryName       *string
	LearningAssessment *int
	Grade              *string
	GroupBased         *string
	Other              *string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

func (s *CourseAssessmentQueryEntity) TableName() string {
	return "course_assessments"
}
