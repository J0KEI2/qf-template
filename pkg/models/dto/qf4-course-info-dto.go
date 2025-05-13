package dto

import (
	"time"

	"github.com/google/uuid"
)

type QF4GetCourseInfoRequestDto struct {
	QF4UID string `query:"uid"`
}

type QF4CreateCourseInfoRequestDto struct {
	QF4ID               uuid.UUID `json:"uid"`
	CategoryName        string    `json:"category_name"`
	CourseCode          string    `json:"course_code"`
	CourseNameTH        string    `json:"course_name_th"`
	CourseNameEN        string    `json:"course_name_en"`
	NumberOfCredits     string    `json:"number_of_credits"`
	CourseTypeID        int       `json:"course_type_id"`
	CourseConditionTH   string    `json:"course_condition_th"`
	CourseConditionEN   string    `json:"course_condition_en"`
	CourseDescriptionTH string    `json:"course_description_th"`
	CourseDescriptionEN string    `json:"course_description_en"`
	CourseObjective     string    `json:"course_objective"`
	StudentActivity     string    `json:"student_activity"`
	FacilitatorTask     string    `json:"facilitator_task"`
	ConsultantTask      string    `json:"consultant_task"`
	StudentGuideline    string    `json:"student_guideline"`
	Location            string    `json:"location"`
	StudentSupport      string    `json:"student_support"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type QF4UpdateCourseInfoRequestDto struct {
	QF4ID               *uuid.UUID `json:"uid"`
	CategoryName        *string    `json:"category_name"`
	CourseCode          *string    `json:"course_code"`
	CourseNameTH        *string    `json:"course_name_th"`
	CourseNameEN        *string    `json:"course_name_en"`
	NumberOfCredits     *string    `json:"number_of_credits"`
	CourseTypeID        *int       `json:"course_type_id"`
	CourseConditionTH   *string    `json:"course_condition_th"`
	CourseConditionEN   *string    `json:"course_condition_en"`
	CourseDescriptionTH *string    `json:"course_description_th"`
	CourseDescriptionEN *string    `json:"course_description_en"`
	CourseObjective     *string    `json:"course_objective"`
	StudentActivity     *string    `json:"student_activity"`
	FacilitatorTask     *string    `json:"facilitator_task"`
	ConsultantTask      *string    `json:"consultant_task"`
	StudentGuideline    *string    `json:"student_guideline"`
	Location            *string    `json:"location"`
	StudentSupport      *string    `json:"student_support"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type QF4CourseInfoResponseDto struct {
	ID                  *int       `json:"id"`
	CategoryName        *string    `json:"category_name"`
	CourseCode          *string    `json:"course_code"`
	CourseNameTH        *string    `json:"course_name_th"`
	CourseNameEN        *string    `json:"course_name_en"`
	NumberOfCredits     *string    `json:"number_of_credits"`
	CourseTypeID        *int       `json:"course_type_id"`
	CourseConditionTH   *string    `json:"course_condition_th"`
	CourseConditionEN   *string    `json:"course_condition_en"`
	CourseDescriptionTH *string    `json:"course_description_th"`
	CourseDescriptionEN *string    `json:"course_description_en"`
	CourseObjective     *string    `json:"course_objective"`
	StudentActivity     *string    `json:"student_activity"`
	FacilitatorTask     *string    `json:"facilitator_task"`
	ConsultantTask      *string    `json:"consultant_task"`
	StudentGuideline    *string    `json:"student_guideline"`
	Location            *string    `json:"location"`
	StudentSupport      *string    `json:"student_support"`
	CreatedAt           *time.Time `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}

type QF4DeleteCourseInfoRequestDto struct {
	QF4UID string `query:"uid"`
}
