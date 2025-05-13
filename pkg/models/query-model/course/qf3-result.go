package query

import "time"

type CourseResult struct {
	ID                  *int       `gorm:"column:id"`
	CategoryName        *string    `gorm:"column:category_name"`
	LearningOutcome     *string    `gorm:"column:learning_outcome"`
	LearningExpectation *string    `gorm:"column:learning_expectation"`
	CreatedAt           *time.Time `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

type CourseResultQueryEntity struct {
	ID                  *int       `gorm:"column:id"`
	CategoryName        *string    `gorm:"column:category_name"`
	LearningOutcome     *string    `gorm:"column:learning_outcome"`
	LearningExpectation *string    `gorm:"column:learning_expectation"`
	CreatedAt           *time.Time `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

func (s *CourseResultQueryEntity) TableName() string {
	return "course_results"
}

func (s CourseResultQueryEntity) String() string {
	return "Result"
}

type CourseResultJointQuery struct {
	ID                  *int       `gorm:"column:id"`
	CategoryName        *string    `gorm:"column:category_name"`
	LearningOutcome     *string    `gorm:"column:learning_outcome"`
	LearningExpectation *string    `gorm:"column:learning_expectation"`
	CreatedAt           *time.Time `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

func (s *CourseResultJointQuery) TableName() string {
	return "course_results"
}

func (s CourseResultJointQuery) String() string {
	return "Result"
}
