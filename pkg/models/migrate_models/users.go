package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	UID                     uuid.UUID      `gorm:"column:uid;type:uuid;primaryKey" json:"uid"`
	SSN                     *string        `gorm:"column:ssn;type:varchar;size:255" json:"ssn,omitempty"`
	NameTH                  string         `gorm:"column:name_th;type:varchar;size:255"`
	NameEN                  string         `gorm:"column:name_en;type:varchar;size:255"`
	Email                   string         `gorm:"column:email;type:varchar;size:255;not null;uniqueIndex" json:"email"`
	Type                    string         `gorm:"column:type;type:user_type;not null" json:"type"`
	FacultyID               uint           `gorm:"column:faculty_id;type:int;default:0" json:"faculty_id"`
	SystemPermissionUID     uuid.UUID      `gorm:"column:system_permission_uid;type:uuid;not null" json:"system_permission_uid"`
	ProgramApprovalMinLevel uint           `gorm:"column:program_approval_min_level;type:smallint;default:0"`
	ProgramApprovalMaxLevel uint           `gorm:"column:program_approval_max_level;type:smallint;default:0"`
	Status                  string         `gorm:"column:status;type:user_status;not null" json:"status"`
	CurrentRoleID           uint           `gorm:"column:current_role_id;type:uint;default:0" json:"current_role_id"`
	CurrentRole             Role           `gorm:"foreignKey:ID;references:CurrentRoleID"`
	LastAccessAt            time.Time      `gorm:"column:last_access_at;type:timestamp" json:"lastAccessAt"`
	CreatedAt               time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt               time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt               gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	// SystemPermission        PermissionSystem    `gorm:"foreignKey:SystemPermissionUID;references:UID" json:"system_permission"`
	CoursePermission  []PermissionCourse  `gorm:"foreignKey:UserUID;references:UID"`
	ProgramPermission []PermissionProgram `gorm:"foreignKey:UserUID;references:UID"`
	MapCoreRoles      []MapCoreRoles      `gorm:"foreignKey:UserID;references:UID"`
	MapFacultiesRoles []MapFacultiesRoles `gorm:"foreignKey:UserID;references:UID"`
	MapProgramRoles   []MapProgramsRoles  `gorm:"foreignKey:UserID;references:UID"`
}
