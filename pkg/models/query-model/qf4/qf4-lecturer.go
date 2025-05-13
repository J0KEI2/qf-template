package query

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

type QF4Lecturer struct {
	ID                int              `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName      string           `gorm:"column:category_name;type:varchar;size:255" json:"category_name"`
	CourseOwnerID     uuid.UUID        `gorm:"column:course_owner_id;type:uuid" json:"course_owner_id"`
	CourseMapLecturer []MapQF4Lecturer `gorm:"foreignKey:QF4LecturerID;references:ID" json:"course_map_lecturer"`
	CreatedAt         time.Time        `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;index" json:"created_at"`
	UpdatedAt         time.Time        `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type QF4LecturerQueryEntity struct {
	ID                *int             `gorm:"column:id"`
	CategoryName      *string          `gorm:"column:category_name"`
	CourseOwner       *query.UserDetail      `gorm:"foreignKey:CourseOwnerID;references:UID"`
	CourseOwnerID     *uuid.UUID       `gorm:"column:course_owner_id"`
	CourseMapLecturer []MapQF4Lecturer `gorm:"foreignKey:QF4LecturerID;references:ID"`
	CreatedAt         *time.Time       `gorm:"column:created_at"`
	UpdatedAt         *time.Time       `gorm:"column:updated_at"`
}

func (s *QF4LecturerQueryEntity) TableName() string {
	return "qf4_lecturers"
}

func (s QF4LecturerQueryEntity) String() string {
	return "Lecturer"
}

type QF4LecturerJoinQuery struct {
	ID                *int                       `gorm:"column:id"`
	CategoryName      *string                    `gorm:"column:category_name"`
	CourseOwner       *query.UserDetail                `gorm:"foreignKey:CourseOwnerID;references:UID"`
	CourseOwnerID     *uuid.UUID                 `gorm:"column:course_owner_id"`
	CourseMapLecturer []MapQF4LecturerJointQuery `gorm:"foreignKey:QF4LecturerID;references:ID"`
	CreatedAt         *time.Time                 `gorm:"column:created_at"`
	UpdatedAt         *time.Time                 `gorm:"column:updated_at"`
}

func (s *QF4LecturerJoinQuery) TableName() string {
	return "qf4_lecturers"
}

func (s QF4LecturerJoinQuery) String() string {
	return "Lecturer"
}

type MapQF4Lecturer struct {
	ID               int        `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	QF4LecturerID    int        `gorm:"column:qf4_lecturer_id;type:integer;size:255;not null" json:"qf4_lecturer_id"`
	CourseLecturer   query.UserDetail `gorm:"foreignKey:CourseLecturerID;references:UID" json:"course_lecturer"`
	CourseLecturerID uuid.UUID  `gorm:"column:course_lecturer_id;type:uuid" json:"course_lecturer_id"`
}

type MapQF4LecturerQueryEntity struct {
	ID               *int        `gorm:"column:id"`
	QF4LecturerID    *int        `gorm:"column:qf4_lecturer_id"`
	CourseLecturer   *query.UserDetail `gorm:"foreignKey:CourseLecturerID;references:UID"`
	CourseLecturerID *uuid.UUID  `gorm:"column:course_lecturer_id"`
}

func (s *MapQF4LecturerQueryEntity) TableName() string {
	return "map_qf4_lecturers"
}

func (s MapQF4LecturerQueryEntity) String() string {
	return "MapQF4Lecturer"
}

type MapQF4LecturerJointQuery struct {
	ID               *int        `gorm:"column:id"`
	QF4LecturerID    *int        `gorm:"column:qf4_lecturer_id"`
	CourseLecturer   *query.UserDetail `gorm:"foreignKey:CourseLecturerID;references:UID"`
	CourseLecturerID *uuid.UUID  `gorm:"column:course_lecturer_id"`
}

func (s *MapQF4LecturerJointQuery) TableName() string {
	return "map_qf4_lecturers"
}

func (s MapQF4LecturerJointQuery) String() string {
	return "MapQF4Lecturer"
}
