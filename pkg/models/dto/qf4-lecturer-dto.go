package dto

import (
	"time"

	"github.com/google/uuid"
)

type QF4GetLecturerRequestDto struct {
	QF4UID string `query:"uid"`
}

type QF4CreateLecturerRequestDto struct {
	QF4ID             uuid.UUID        `json:"uid"`
	CategoryName      string           `json:"category_name"`
	CourseOwnerID     uuid.UUID        `json:"course_owner_id"`
	CourseMapLecturer []MapQF4Lecturer `json:"course_map_lecturer"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

type MapQF4Lecturer struct {
	ID               int       `json:"id"`
	QF4LecturerID    int       `json:"qf4_lecturer_id"`
	CourseLecturerID uuid.UUID `json:"course_lecturer_id"`
}

type QF4UpdateLecturerRequestDto struct {
	QF4ID              *uuid.UUID       `json:"uid"`
	CategoryName       *string          `json:"category_name"`
	CourseOwnerID      *uuid.UUID       `json:"course_owner_id"`
	CourseMapLecturers []MapQF4Lecturer `json:"course_map_lecturers"`
	CreatedAt          *time.Time       `json:"created_at"`
	UpdatedAt          *time.Time       `json:"updated_at"`
}

type QF4LecturerResponseDto struct {
	ID                 *int             `json:"id"`
	CategoryName       *string          `json:"category_name"`
	CourseOwnerID      *uuid.UUID       `json:"course_owner_id"`
	CourseMapLecturers []MapQF4Lecturer `json:"course_map_lecturers"`
	CreatedAt          *time.Time       `json:"created_at"`
	UpdatedAt          *time.Time       `json:"updated_at"`
}
