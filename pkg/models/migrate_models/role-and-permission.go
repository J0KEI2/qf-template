package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MapCoreRoles struct {
	ID        uint           `gorm:"column:id;primaryKey;type:bigint" json:"id"`
	UserID    uuid.UUID      `gorm:"column:user_id;type:uuid" json:"user_id"`
	RoleID    uint           `gorm:"column:role_id;type:bigint" json:"role_id"`
	CreatedAt time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (MapCoreRoles) TableName() string {
	return "map_core_roles"
}

type MapFacultiesRoles struct {
	ID        uint           `gorm:"column:id;primaryKey;type:bigint" json:"id"`
	UserID    uuid.UUID      `gorm:"column:user_id;type:uuid" json:"user_id"`
	RoleID    uint           `gorm:"column:role_id;type:bigint" json:"role_id"`
	FacultyID uint           `gorm:"column:faculty_id;type:bigint" json:"faculty_id"`
	CreatedAt time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (MapFacultiesRoles) TableName() string {
	return "map_faculties_roles"
}

type MapProgramsRoles struct {
	ID        uint           `gorm:"column:id;primaryKey;type:bigint" json:"id"`
	UserID    uuid.UUID      `gorm:"column:user_id;type:uuid" json:"user_id"`
	RoleID    uint           `gorm:"column:role_id;type:bigint" json:"role_id"`
	ProgramID uuid.UUID      `gorm:"column:program_id;type:uuid" json:"program_id"`
	CreatedAt time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (MapProgramsRoles) TableName() string {
	return "map_programs_roles"
}

type Role struct {
	ID                   uint                `gorm:"column:id;primaryKey;type:bigint" json:"id"`
	RoleNameTH           string              `gorm:"column:role_name_th;type:varchar(255)" json:"role_name_th"`
	RoleNameEN           string              `gorm:"column:role_name_en;type:varchar(255)" json:"role_name_en"`
	ProgramRoleType      uint                `gorm:"column:program_role_type;type:bigint;default:0" json:"program_role_type"`
	CourseRoleType       uint                `gorm:"column:course_role_type;type:bigint;default:0" json:"course_role_type"`
	SystemLevel          uint                `gorm:"column:system_level;type:bigint;default:0"`
	ProgramApprovalLevel uint                `gorm:"column:program_approval_level;type:bigint;default:0" json:"program_approval_level"`
	CourseApprovalLevel  uint                `gorm:"column:course_approval_level;type:bigint;default:0" json:"course_approval_level"`
	ProgramActionLevel   uint                `gorm:"column:program_action_level;type:bigint;default:0" json:"program_action_level"`
	CourseActionLevel    uint                `gorm:"column:course_action_level;type:bigint;default:0" json:"course_action_level"`
	ProgramAccessLevel   uint                `gorm:"column:program_access_level;type:bigint;default:0" json:"program_access_level"`
	CourseAccessLevel    uint                `gorm:"column:course_access_level;type:bigint;default:0" json:"course_access_level"`
	CreatedAt            time.Time           `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt            time.Time           `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	MapCoreRoles         []MapCoreRoles      `gorm:"foreignKey:RoleID;references:ID"`
	MapFacultiesRoles    []MapFacultiesRoles `gorm:"foreignKey:RoleID;references:ID"`
	MapProgramRoles      []MapProgramsRoles  `gorm:"foreignKey:RoleID;references:ID"`
	DeletedAt            gorm.DeletedAt      `gorm:"index;column:deleted_at"`
}

func (Role) TableName() string {
	return "roles"
}
