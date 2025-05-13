package query

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

type CourseLecturer struct {
	ID                int                 `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName      string              `gorm:"column:category_name;type:varchar;size:255" json:"category_name"`
	CourseOwnerID     uuid.UUID           `gorm:"column:course_owner_id;type:uuid" json:"course_owner_id"`
	CourseMapLecturer []MapCourseLecturer `gorm:"foreignKey:CourseLecturerID;references:ID" json:"course_map_lecturer"`
	CreatedAt         time.Time           `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;index" json:"created_at"`
	UpdatedAt         time.Time           `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CourseLecturerQueryEntity struct {
	ID                *int                `gorm:"column:id"`
	CategoryName      *string             `gorm:"column:category_name"`
	CourseOwner       *query.UserDetail   `gorm:"foreignKey:CourseOwnerID;references:UID"`
	CourseOwnerID     *uuid.UUID          `gorm:"column:course_owner_id"`
	CourseMapLecturer []MapCourseLecturer `gorm:"foreignKey:CourseLecturerID;references:ID"`
	CreatedAt         *time.Time          `gorm:"column:created_at"`
	UpdatedAt         *time.Time          `gorm:"column:updated_at"`
}

func (s *CourseLecturerQueryEntity) TableName() string {
	return "course_lecturers"
}

func (s CourseLecturerQueryEntity) String() string {
	return "Lecturer"
}

type CourseLecturerJoinQuery struct {
	ID                *int                          `gorm:"column:id"`
	CategoryName      *string                       `gorm:"column:category_name"`
	CourseOwner       *query.UserDetail             `gorm:"foreignKey:CourseOwnerID;references:UID"`
	CourseOwnerID     *uuid.UUID                    `gorm:"column:course_owner_id"`
	CourseMapLecturer []MapCourseLecturerJointQuery `gorm:"foreignKey:CourseLecturerID;references:ID"`
	CreatedAt         *time.Time                    `gorm:"column:created_at"`
	UpdatedAt         *time.Time                    `gorm:"column:updated_at"`
}

func (s *CourseLecturerJoinQuery) TableName() string {
	return "course_lecturers"
}

func (s CourseLecturerJoinQuery) String() string {
	return "Lecturer"
}

type MapCourseLecturer struct {
	ID               int               `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CourseLecturerID int               `gorm:"column:course_lecturer_id;type:integer;size:255;not null" json:"course_lecturer_id"`
	CourseLecturer   *query.UserDetail `gorm:"foreignKey:UID;references:EmployeeID"`
	EmployeeID       *uuid.UUID        `gorm:"column:employee_id"`
}

type MapCourseLecturerQueryEntity struct {
	ID               *int              `gorm:"column:id"`
	CourseLecturerID *int              `gorm:"column:course_lecturer_id"`
	CourseLecturer   *query.UserDetail `gorm:"foreignKey:UID;references:EmployeeID"`
	EmployeeID       *uuid.UUID        `gorm:"column:employee_id"`
}

func (s *MapCourseLecturerQueryEntity) TableName() string {
	return "map_course_lecturers"
}

func (s MapCourseLecturerQueryEntity) String() string {
	return "MapCourseLecturer"
}

type MapCourseLecturerJointQuery struct {
	ID               *int              `gorm:"column:id"`
	CourseLecturerID *int              `gorm:"column:course_lecturer_id"`
	CourseLecturer   *query.UserDetail `gorm:"foreignKey:UID;references:EmployeeID"`
	EmployeeID       *uuid.UUID        `gorm:"column:employee_id"`
}

func (s *MapCourseLecturerJointQuery) TableName() string {
	return "map_course_lecturers"
}

func (s MapCourseLecturerJointQuery) String() string {
	return "MapCourseLecturer"
}
