package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramYearAndSemesterQueryEntity struct {
	ID               *uint                            `gorm:"column:id"`
	ProgramSubPlan   *ProgramSubPlanQueryEntity       `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID *uint                            `gorm:"column:program_sub_plan_id;type:uint"`
	CourseDetail     []ProgramCourseDetailQueryEntity `gorm:"foreignKey:YearAndSemesterID;references:ID"`
	Year             *string                          `gorm:"column:year"`
	Semester         *string                          `gorm:"column:semester"`
	CreatedAt        *time.Time                       `gorm:"column:created_at"`
	UpdatedAt        *time.Time                       `gorm:"column:updated_at"`
	DeletedAt        *gorm.DeletedAt                  `gorm:"column:deleted_at"`
}

func (ProgramYearAndSemesterQueryEntity) TableName() string {
	return "program_year_and_semester"
}
