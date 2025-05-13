package query

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionSystemQueryEntity struct {
	UID                       *uuid.UUID      `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	RoleNameTH                *string         `gorm:"column:role_name_th;type:varchar;size:255"`
	RoleNameEN                *string         `gorm:"column:role_name_en;type:varchar;size:255"`
	PageAccessibility         *string         `gorm:"column:page_accessibility;type:text"`
	ProgramAccessibility      *string         `gorm:"column:program_accessibility;type:text"`
	CourseAccessibility       *string         `gorm:"column:course_accessibility;type:text"`
	ProgramAccessibilityLevel *uint           `gorm:"column:program_accessibility_level;type:smallint"`
	CourseAccessibilityLevel  *uint           `gorm:"column:course_accessibility_level;type:smallint"`
	UAMControl                bool            `gorm:"column:uam_control;type:tinyint"`
	CanComment                bool            `gorm:"column:can_comment;type:tinyint"`
	CanApproved               bool            `gorm:"column:can_approved;type:tinyint"`
	CreatedAt                 *time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                 *time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                 *gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (PermissionSystemQueryEntity) TableName() string {
	return "permission_system"
}

type PermissionProgramQueryEntity struct {
	ID            *int            `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	UserUID       *uuid.UUID      `gorm:"column:user_uid;type:uuid"`
	ProgramUID    *uuid.UUID      `gorm:"column:program_uid;type:uuid"`
	Accessibility *uint           `gorm:"column:accessibility;type:smallint"`
	CreatedAt     *time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt     *gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (PermissionProgramQueryEntity) TableName() string {
	return "permission_program"
}

type PermissionCourseQueryEntity struct {
	ID            *int            `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	UserUID       *uuid.UUID      `gorm:"column:user_uid;type:uuid"`
	CourseUID     *uuid.UUID      `gorm:"column:course_uid;type:uuid;not null"`
	Accessibility *int            `gorm:"column:accessibility;type:smallint"`
	CreatedAt     *time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt     *gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (PermissionCourseQueryEntity) TableName() string {
	return "permission_course"
}
