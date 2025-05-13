package dto

import (
	"time"

	"github.com/google/uuid"
)

type CourseGetCourseInfoRequestDto struct {
	CourseUID string `query:"uid"`
}

type CourseCreateCourseInfoRequestDto struct {
	CourseID               uuid.UUID `json:"uid"`
	CategoryName        *string   `json:"category_name"`
	CourseCode          *string   `json:"course_code"`
	CourseNameTH        *string   `json:"course_name_th"`
	CourseNameEN        *string   `json:"course_name_en"`
	TotalCredit         *uint     `json:"total_credit"`
	Credit1             *uint     `json:"credit_1"`
	Credit2             *uint     `json:"credit_2"`
	Credit3             *uint     `json:"credit_3"`
	CourseTypeID        *int      `json:"course_type_id"`
	CourseConditionTH   *string   `json:"course_condition_th"`
	CourseConditionEN   *string   `json:"course_condition_en"`
	CourseDescriptionTH *string   `json:"course_description_th"`
	CourseDescriptionEN *string   `json:"course_description_en"`
	CourseObjective     *string   `json:"course_objective"`
	Location            *string   `json:"location"`
}

type CourseUpdateCourseInfoRequestDto struct {
	CourseID               *uuid.UUID `json:"uid"`
	CategoryName        *string    `json:"category_name"`
	CourseCode          *string    `json:"course_code"`
	CourseNameTH        *string    `json:"course_name_th"`
	CourseNameEN        *string    `json:"course_name_en"`
	TotalCredit         *uint     `json:"total_credit"`
	Credit1             *uint     `json:"credit_1"`
	Credit2             *uint     `json:"credit_2"`
	Credit3             *uint     `json:"credit_3"`
	CourseTypeID        *int       `json:"course_type_id"`
	CourseConditionTH   *string    `json:"course_condition_th"`
	CourseConditionEN   *string    `json:"course_condition_en"`
	CourseDescriptionTH *string    `json:"course_description_th"`
	CourseDescriptionEN *string    `json:"course_description_en"`
	CourseObjective     *string    `json:"course_objective"`
	Location            *string    `json:"location"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type CourseInfoResponseDto struct {
	ID                  *int       `json:"id"`
	CategoryName        *string    `json:"category_name"`
	CourseCode          *string    `json:"course_code"`
	CourseNameTH        *string    `json:"course_name_th"`
	CourseNameEN        *string    `json:"course_name_en"`
	TotalCredit         *uint      `json:"total_credit"`
	Credit1             *uint      `json:"credit_1"`
	Credit2             *uint      `json:"credit_2"`
	Credit3             *uint      `json:"credit_3"`
	CourseTypeID        *int       `json:"course_type_id"`
	CourseConditionTH   *string    `json:"course_condition_th"`
	CourseConditionEN   *string    `json:"course_condition_en"`
	CourseDescriptionTH *string    `json:"course_description_th"`
	CourseDescriptionEN *string    `json:"course_description_en"`
	CourseObjective     *string    `json:"course_objective"`
	Location            *string    `json:"location"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type CourseDeleteCourseInfoRequestDto struct {
	CourseUID string `query:"uid"`
}
