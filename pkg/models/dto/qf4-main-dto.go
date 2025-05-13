package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
)

type QF4GetMainRequestDto struct {
	QF4UID string `query:"uid"`
}

type QF4CreateMainRequestDto struct {
	QF4ID                        uuid.UUID        `json:"qf4_id"`
	CourseNumber                 int              `json:"course_number"`
	Version                      string           `json:"version"`
	FacultyID                    uint             `json:"faculty_id"`
	DepartmentName               string           `json:"department_name"`
	EducationYear                string           `json:"education_year"`
	QF4CourseInfoID              *int             `json:"qf4_course_info_id"`
	QF4LecturerID                *int             `json:"qf4_lecturer_id"`
	QF4ResultID                  *int             `json:"qf4_result_id"`
	QF4CourseTypeAndManagementID *int             `json:"qf4_course_type_and_management_id"`
	QF4AssessmentID              *int             `json:"qf4_assessment_id"`
	QF4ReferenceID               *int             `json:"qf4_reference_id"`
	Status                       enums.UserStatus `json:"status"`
	CreatedAt                    time.Time        `json:"created_at"`
	UpdatedAt                    time.Time        `json:"updated_at"`
}

type QF4UpdateMainRequestDto struct {
	QF4ID                        *uuid.UUID        `json:"qf4_id"`
	CourseNumber                 *int              `json:"course_number"`
	Version                      *string           `json:"version"`
	FacultyID                    *uint             `json:"faculty_id"`
	DepartmentName               *string           `json:"department_name"`
	EducationYear                *string           `json:"education_year"`
	QF4CourseInfoID              *int              `json:"qf4_course_info_id"`
	QF4LecturerID                *int              `json:"qf4_lecturer_id"`
	QF4ResultID                  *int              `json:"qf4_result_id"`
	QF4CourseTypeAndManagementID *int              `json:"qf4_course_type_and_management_id"`
	QF4AssessmentID              *int              `json:"qf4_assessment_id"`
	QF4ReferenceID               *int              `json:"qf4_reference_id"`
	Status                       *enums.UserStatus `json:"status"`
	CreatedAt                    *time.Time        `json:"created_at"`
	UpdatedAt                    *time.Time        `json:"updated_at"`
}

type QF4MainResponseDto struct {
	QF4ID                        *uuid.UUID        `json:"qf4_id"`
	CourseNumber                 *int              `json:"course_number"`
	Version                      *string           `json:"version"`
	FacultyID                    *uint             `json:"faculty_id"`
	DepartmentName               *string           `json:"department_name"`
	EducationYear                *string           `json:"education_year"`
	QF4CourseInfoID              *int              `json:"qf4_course_info_id"`
	QF4LecturerID                *int              `json:"qf4_lecturer_id"`
	QF4ResultID                  *int              `json:"qf4_result_id"`
	QF4CourseTypeAndManagementID *int              `json:"qf4_course_type_and_management_id"`
	QF4AssessmentID              *int              `json:"qf4_assessment_id"`
	QF4ReferenceID               *int              `json:"qf4_reference_id"`
	Status                       *enums.UserStatus `json:"status"`
	CreatedAt                    *time.Time        `json:"created_at"`
	UpdatedAt                    *time.Time        `json:"updated_at"`
}

type QF4DeleteMainRequestDto struct {
	QF4UID string `query:"uid"`
}
