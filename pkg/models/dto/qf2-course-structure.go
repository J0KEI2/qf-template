package dto

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type GetCourseStructureRequestDto struct {
	PlanDetailID int `json:"plan_detail_id"`
	ProgramId    int `json:"program_id"`
}

type GetCoursePaginationForStructureResponseDto struct {
	Items []CoursePaginationForStructure `json:"items"`
	*models.PaginationOptions
}

type CreateOrUpdateCourseStructureRequestDto struct {
	Items []ProgramStructureRequestDto `json:"items"`
}

type GetCourseStructureResponseDto struct {
	Items []ProgramStructureResponseDto `json:"items"`
	*models.PaginationOptions
}

type GetCourseByStructureResponseDto struct {
	Items []ProgramCourseDetailResponseDto `json:"items"`
	*models.PaginationOptions
}

type CoursePaginationForStructure struct {
	ID                  *int       `json:"id"`
	Name                *string    `json:"name"`
	CategoryName        *string    `json:"category_name"`
	CourseCode          *string    `json:"course_code"`
	CourseYear          *string    `json:"course_year"`
	CourseSource        *string    `json:"course_source"`
	CourseNameTH        *string    `json:"course_name_th"`
	CourseNameEN        *string    `json:"course_name_en"`
	Version             *string    `json:"course_version"`
	CourseTypeID        *int       `json:"course_type_id"`
	CourseCredit        *uint      `json:"course_credit"`
	CreditDetail        *string    `json:"credit_detail"`
	Credit1             *uint      `json:"credit_1"`
	Credit2             *uint      `json:"credit_2"`
	Credit3             *uint      `json:"credit_3"`
	CourseConditionTH   *string    `json:"course_condition_th"`
	CourseConditionEN   *string    `json:"course_condition_en"`
	CourseDescriptionTH *string    `json:"course_description_th"`
	CourseDescriptionEN *string    `json:"course_description_en"`
	CourseObjective     *string    `json:"course_objective"`
	REGKkuKey           *string    `json:"reg_kku_key"`
	CourseKey           *uuid.UUID `json:"course_key"`
	Location            *string    `json:"location"`
	IsEditedCourse      *bool      `json:"is_edited_course"`
	IsNewCourse         *bool      `json:"is_new_course"`
}
