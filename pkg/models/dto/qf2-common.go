package dto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramStructureRequestDto struct {
	ID              *uint                           `json:"id"`
	SubPlanID       *uint                           `json:"sub_plan_id"`
	PlanDetailID    *uint                           `json:"plan_detail_id"`
	Name            *string                         `json:"name"`
	Order           *uint                           `json:"order"`
	ParentID        *uint                           `json:"parent_id"`
	Children        []ProgramStructureRequestDto    `json:"children"`
	CourseDetails   []ProgramCourseDetailRequestDto `json:"course_details"`
	Qualification   *string                         `json:"qualification"`
	StructureCredit *json.Number                    `json:"credit"`
	CreatedAt       *time.Time                      `json:"created_at"`
	UpdatedAt       *time.Time                      `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt                 `json:"deleted_at"`
}

type ProgramStructureResponseDto struct {
	ID               *uint                            `json:"id"`
	ProgramSubPlanID *uint                            `json:"program_sub_plan_id"`
	Name             *string                          `json:"name"`
	Order            *uint                            `json:"order"`
	ParentID         *uint                            `json:"parent_id"`
	Children         []ProgramStructureResponseDto    `json:"children"`
	CourseDetails    []ProgramCourseDetailResponseDto `json:"course_details"`
	Qualification    *string                          `json:"qualification"`
	StructureCredit  *uint                            `json:"credit"`
	CreatedAt        *time.Time                       `json:"created_at"`
	UpdatedAt        *time.Time                       `json:"updated_at"`
	DeletedAt        *gorm.DeletedAt                  `json:"deleted_at"`
}

type ProgramCourseDetailRequestDto struct {
	ID                  *uint        `json:"id"`
	ProgramStructureID  *uint        `json:"program_structure_id"`
	ProgramSubPlanID    *uint        `json:"program_sub_plan_id"`
	YearAndSemesterID   *uint        `json:"year_and_semester_id"`
	CourseTypeID        *int         `json:"course_type_id"`
	CourseType          *string      `json:"course_type"`
	CourseCode          *string      `json:"course_code"`
	CourseYear          *string      `json:"course_year"`
	CourseSource        *string      `json:"course_source"`
	REGKkuKey           *string      `json:"reg_kku_key"`
	CourseKey           *uuid.UUID   `json:"course_key"`
	CourseNameTH        *string      `json:"course_name_th"`
	CourseNameEN        *string      `json:"course_name_en"`
	Version             *string      `json:"course_version"`
	CourseCredit        *json.Number `json:"course_credit"`
	Credit1             *uint        `json:"credit_1"`
	Credit2             *uint        `json:"credit_2"`
	Credit3             *uint        `json:"credit_3"`
	CourseConditionTH   *string      `json:"course_condition_th"`
	CourseConditionEN   *string      `json:"course_condition_en"`
	CourseDescriptionEN *string      `json:"course_description_en"`
	CourseDescriptionTH *string      `json:"course_description_th"`
	CourseObjective     *string      `json:"course_objective"`
	IsCreditCalc        *bool        `json:"is_credit_calc"`
	IsEditedCourse      *bool        `json:"is_edited_course"`
	IsNewCourse         *bool        `json:"is_new_course"`
}

type ProgramCourseDetailResponseDto struct {
	ID                *uint           `json:"id"`
	YearAndSemesterID *uint           `json:"year_and_semester_id"`
	CourseTypeID      *int            `json:"course_type_id"`
	CourseType        *string         `json:"course_type"`
	CourseCode        *string         `json:"course_code"`
	CourseYear        *string         `json:"course_year"`
	CourseNameTH      *string         `json:"course_name_th"`
	CourseNameEN      *string         `json:"course_name_en"`
	Name              *string         `json:"name"`
	CourseCredit      *uint           `json:"credit"`
	CreditDetail      *string         `json:"credit_detail"`
	Credit1           *uint           `json:"credit_1"`
	Credit2           *uint           `json:"credit_2"`
	Credit3           *uint           `json:"credit_3"`
	IsCreditCalc      *bool           `json:"is_credit_calc"`
	IsEditedCourse    *bool           `json:"is_edited_course"`
	IsNewCourse       *bool           `json:"is_new_course"`
	CreatedAt         *time.Time      `json:"created_at"`
	UpdatedAt         *time.Time      `json:"updated_at"`
	DeletedAt         *gorm.DeletedAt `json:"deleted_at"`
}

type ProgramYearRequestDto struct {
	Year      *string                     `json:"year"`
	Semesters []ProgramSemesterRequestDto `json:"semesters"`
}

type ProgramSemesterRequestDto struct {
	ID            *uint                           `json:"id"`
	PlanDetailID  *uint                           `json:"plan_detail_id"`
	CourseDetails []ProgramCourseDetailRequestDto `json:"course_details"`
	Semester      *string                         `json:"semester"`
	CreatedAt     *time.Time                      `json:"created_at"`
	UpdatedAt     *time.Time                      `json:"updated_at"`
	DeletedAt     *gorm.DeletedAt                 `json:"deleted_at"`
}

type ProgramYearAndSemesterCourseRequestDto struct {
	CourseID *uint `json:"course_id"`
}

type ProgramYearAndSemesterRequestDto struct {
	ID        *uint           `json:"id"`
	SubPlanID *uint           `json:"sub_plan_id"`
	Year      *string         `json:"year"`
	Semester  *string         `json:"semester"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

type ProgramYearAndSemesterResponseDto struct {
	ID            *uint                            `json:"id"`
	SubPlanID     *uint                            `json:"sub_plan_id"`
	Year          *string                          `json:"year"`
	Semester      *string                          `json:"semester"`
	CourseDetails []ProgramCourseDetailResponseDto `json:"course_details,omitempty"`
	CreatedAt     *time.Time                       `json:"created_at"`
	UpdatedAt     *time.Time                       `json:"updated_at"`
	DeletedAt     *gorm.DeletedAt                  `json:"deleted_at"`
}

type ProgramYearResponseDto struct {
	Year      *string                      `json:"year"`
	Semesters []ProgramSemesterResponseDto `json:"semesters"`
}

type ProgramSemesterResponseDto struct {
	ID            *uint                            `json:"id"`
	SubPlanID     *uint                            `json:"sub_plan_id"`
	CourseDetails []ProgramCourseDetailResponseDto `json:"course_details"`
	Semester      *string                          `json:"semester"`
	CreatedAt     *time.Time                       `json:"created_at"`
	UpdatedAt     *time.Time                       `json:"updated_at"`
	DeletedAt     *gorm.DeletedAt                  `json:"deleted_at"`
}
