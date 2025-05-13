package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Faculty struct {
	ID                uint                `gorm:"column:id;primaryKey;autoIncrement;index;uniqueIndex"`
	FacultyID         string              `gorm:"column:faculty_id;type:varchar;size:255;not null" json:"facultyID,omitempty"`
	University        string              `gorm:"column:university;type:varchar;size:255;not null" json:"university,omitempty"`
	FacultyNameTH     string              `gorm:"column:faculty_name_th;type:varchar;size:255;not null" json:"facultyNameTH,omitempty"`
	FacultyNameEN     string              `gorm:"column:faculty_name_en;type:varchar;size:255;not null" json:"facultyNameEN,omitempty"`
	MapFacultiesRoles []MapFacultiesRoles `gorm:"foreignKey:FacultyID;references:ID"`
	CreatedAt         time.Time           `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time           `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt         gorm.DeletedAt      `gorm:"column:deleted_at;type:timestamp"`
}

func (Faculty) TableName() string {
	return "common_faculties"
}

type CourseType struct {
	ID             int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CourseTypeName string         `gorm:"column:course_type_name;type:varchar;size:255;not null" json:"courseTypeName,omitempty"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type EmployeeDetails struct {
	UID                 uuid.UUID      `gorm:"column:uid;type:uuid;not null;primaryKey"`
	Email               string         `gorm:"column:email;type:varchar;size:255;not null;uniqueIndex"`
	TitleTh             string         `gorm:"column:title_th;type:varchar;size:255;not null"`
	TitleEn             string         `gorm:"column:title_en;type:varchar;size:255;not null"`
	FirstnameTh         string         `gorm:"column:firstname_th;type:varchar;size:255;not null"`
	FirstnameEn         string         `gorm:"column:firstname_en;type:varchar;size:255;not null"`
	LastnameTh          string         `gorm:"column:lastname_th;type:varchar;size:255;not null"`
	LastnameEn          string         `gorm:"column:lastname_en;type:varchar;size:255;not null"`
	EducationBackGround *string        `gorm:"column:education_back_ground;type:varchar;size:511"`
	Position            *string        `gorm:"column:position;type:varchar"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;index"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;"`
}

func (EmployeeDetails) TableName() string {
	return "common_employee_details"
}

type ReferenceOption struct {
	ID        int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	Name      string         `gorm:"column:name;type:varchar;size:255;not null" json:"name,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

func (ReferenceOption) TableName() string {
	return "common_reference_options"
}
