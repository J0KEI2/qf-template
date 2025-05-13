package query

import (
	"time"

	"github.com/google/uuid"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
	"gorm.io/gorm"
)

type ProgramCourseDetailQueryEntity struct {
	ID                  *uint                        `gorm:"column:id"`
	YearAndSemesterID   *uint                        `gorm:"column:year_and_semester_id"`
	ProgramStructureID  *uint                        `gorm:"column:program_structure_id"`
	ProgramSubPlanID    *uint                        `gorm:"column:program_sub_plan_id"`
	CourseSource        *string                      `gorm:"column:course_source;"` // REG || Course
	REGKkuKey           *string                      `gorm:"column:reg_kku_key;"`
	CourseKey           *uuid.UUID                   `gorm:"column:course_key;"`
	CourseMain          *query.CourseMainQueryEntity `gorm:"foreignKey:CourseID;references:CourseKey"`
	CourseTypeID        *int                         `gorm:"column:course_type_id"`
	CourseType          *string                      `gorm:"column:course_type"`
	CourseCode          *string                      `gorm:"column:course_code"`
	CourseYear          *string                      `gorm:"column:course_year"`
	CourseNameTH        *string                      `gorm:"column:course_name_th"`
	CourseNameEN        *string                      `gorm:"column:course_name_en"`
	Version             *string                      `gorm:"column:course_version"`
	CourseCredit        *uint                        `gorm:"column:course_credit"`
	Credit1             *uint                        `gorm:"column:credit_1;type:int;not null;default:0"`
	Credit2             *uint                        `gorm:"column:credit_2;type:int;not null;default:0"`
	Credit3             *uint                        `gorm:"column:credit_3;type:int;not null;default:0"`
	CourseConditionTH   *string                      `gorm:"column:course_condition_th;type:varchar;size:255"`
	CourseConditionEN   *string                      `gorm:"column:course_condition_en;type:varchar;size:255"`
	CourseDescriptionEN *string                      `gorm:"column:course_description_th;type:varchar;size:255"`
	CourseDescriptionTH *string                      `gorm:"column:course_description_en;type:varchar;size:255"`
	CourseObjective     *string                      `gorm:"column:course_objective;type:varchar;size:255"`
	IsCreditCalc        *bool                        `gorm:"column:is_credit_calc"`
	IsEditedCourse      *bool                        `gorm:"column:is_edited_course"`
	IsNewCourse         *bool                        `gorm:"column:is_new_course"`
	CreatedAt           *time.Time                   `gorm:"column:created_at"`
	UpdatedAt           *time.Time                   `gorm:"column:updated_at"`
	DeletedAt           *gorm.DeletedAt              `gorm:"column:deleted_at"`
}

func (q *ProgramCourseDetailQueryEntity) TableName() string {
	return "program_course_detail"
}
