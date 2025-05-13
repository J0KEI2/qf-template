package query

import "time"

type QF4Result struct {
	ID                  *int       `gorm:"column:id"`
	CategoryName        *string    `gorm:"column:category_name"`
	LearningOutcome     *string    `gorm:"column:learning_outcome"`
	LearningExpectation *string    `gorm:"column:learning_expectation"`
	CreatedAt           *time.Time `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

type QF4ResultQueryEntity struct {
	ID                  *int       `gorm:"column:id"`
	CategoryName        *string    `gorm:"column:category_name"`
	LearningOutcome     *string    `gorm:"column:learning_outcome"`
	LearningExpectation *string    `gorm:"column:learning_expectation"`
	CreatedAt           *time.Time `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

func (s *QF4ResultQueryEntity) TableName() string {
	return "qf4_results"
}

func (s QF4ResultQueryEntity) String() string {
	return "Result"
}

type QF4ResultJointQuery struct {
	ID                  *int       `gorm:"column:id"`
	CategoryName        *string    `gorm:"column:category_name"`
	LearningOutcome     *string    `gorm:"column:learning_outcome"`
	LearningExpectation *string    `gorm:"column:learning_expectation"`
	CreatedAt           *time.Time `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

func (s *QF4ResultJointQuery) TableName() string {
	return "qf4_results"
}

func (s QF4ResultJointQuery) String() string {
	return "Result"
}
