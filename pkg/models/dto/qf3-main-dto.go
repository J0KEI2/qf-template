package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type CourseGetMainRequestDto struct {
	CourseUID string `query:"uid"`
}

type CourseCreateMainRequestDto struct {
	CourseID                  uuid.UUID `json:"course_id"`
	CourseNumber              int       `json:"course_number"`
	Version                   string    `json:"version"`
	FacultyID                 uint      `json:"faculty_id"`
	DepartmentName            string    `json:"department_name"`
	EducationYear             string    `json:"education_year"`
	CourseInfoID              *int      `json:"course_info_id"`
	CourseLecturerID          *int      `json:"course_lecturer_id"`
	CourseResultID            *int      `json:"course_result_id"`
	CourseTypeAndManagementID *int      `json:"course_type_and_management_id"`
	CourseAssessmentID        *int      `json:"course_assessment_id"`
	CourseReferenceID         *int      `json:"course_reference_id"`
	Status                    string    `json:"status"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}

type CourseUpdateMainRequestDto struct {
	CourseID                  *uuid.UUID `json:"course_id"`
	CourseNumber              *int       `json:"course_number"`
	Version                   *string    `json:"version"`
	FacultyID                 *uint      `json:"faculty_id"`
	DepartmentName            *string    `json:"department_name"`
	EducationYear             *string    `json:"education_year"`
	CourseInfoID              *int       `json:"course_info_id"`
	CourseLecturerID          *int       `json:"course_lecturer_id"`
	CourseResultID            *int       `json:"course_result_id"`
	CourseTypeAndManagementID *int       `json:"course_type_and_management_id"`
	CourseAssessmentID        *int       `json:"course_assessment_id"`
	CourseReferenceID         *int       `json:"course_reference_id"`
	Status                    *string    `json:"status"`
	CreatedAt                 *time.Time `json:"created_at"`
	UpdatedAt                 *time.Time `json:"updated_at"`
}

type GetMainCoursePaginationResponseDto struct {
	Items []CourseMainWithInfoResponseDto `json:"items"`
	models.PaginationOptions
}

type CourseMainWithInfoResponseDto struct {
	CourseID       *uuid.UUID `json:"course_id"`
	CourseNumber   *int       `json:"course_number"`
	Version        *string    `json:"version"`
	DepartmentName *string    `json:"department_name"`
	EducationYear  *string    `json:"education_year"`
	Status         *string    `json:"status"`
	CourseCode     *string    `json:"course_code"`
	CourseNameTh   *string    `json:"course_name_th"`
	CourseNameEn   *string    `json:"course_name_en"`
	CategoryName   *string    `json:"category_name"`
	TotalCredit    *uint      `json:"total_credit"`
	Credit1        *uint      `json:"credit1"`
	Credit2        *uint      `json:"credit2"`
	Credit3        *uint      `json:"credit3"`
}

type CourseMainResponseDto struct {
	CourseID                  *uuid.UUID `json:"course_id"`
	CourseNumber              *int       `json:"course_number"`
	Version                   *string    `json:"version"`
	FacultyID                 *uint      `json:"faculty_id"`
	DepartmentName            *string    `json:"department_name"`
	EducationYear             *string    `json:"education_year"`
	CourseInfoID              *int       `json:"course_info_id"`
	CourseLecturerID          *int       `json:"course_lecturer_id"`
	CourseResultID            *int       `json:"course_result_id"`
	CourseTypeAndManagementID *int       `json:"course_type_and_management_id"`
	CourseAssessmentID        *int       `json:"course_assessment_id"`
	CourseReferenceID         *int       `json:"course_reference_id"`
	Status                    *string    `json:"status"`
	CreatedAt                 *time.Time `json:"created_at"`
	UpdatedAt                 *time.Time `json:"updated_at"`
}

type CourseDeleteMainRequestDto struct {
	CourseUID string `query:"uid"`
}
