package dto

import (
	"time"

	"github.com/google/uuid"
)

type CourseGetLecturerRequestDto struct {
	CourseUID string `query:"uid"`
}

type CourseCreateLecturerRequestDto struct {
	CourseID           uuid.UUID           `json:"uid"`
	CategoryName       string              `json:"category_name"`
	CourseOwnerID      uuid.UUID           `json:"course_owner_id"`
	CourseMapLecturers []MapCourseLecturer `json:"course_map_lecturers"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

type MapCourseLecturer struct {
	ID               int       `json:"id"`
	CourseLecturerID int       `json:"course_lecturer_id"`
	EmployeeID       uuid.UUID `json:"employee_id"`
}

type CourseUpdateLecturerRequestDto struct {
	CourseID           *uuid.UUID          `json:"uid"`
	CategoryName       *string             `json:"category_name"`
	CourseOwnerID      *uuid.UUID          `json:"course_owner_id"`
	CourseMapLecturers []MapCourseLecturer `json:"course_map_lecturers"`
	CreatedAt          *time.Time          `json:"created_at"`
	UpdatedAt          *time.Time          `json:"updated_at"`
}

type CourseLecturerResponseDto struct {
	ID                 *int                `json:"id"`
	CategoryName       *string             `json:"category_name"`
	CourseOwnerID      *uuid.UUID          `json:"course_owner_id"`
	CourseMapLecturers []MapCourseLecturer `json:"course_map_lecturers"`
	CreatedAt          *time.Time          `json:"created_at"`
	UpdatedAt          *time.Time          `json:"updated_at"`
}
