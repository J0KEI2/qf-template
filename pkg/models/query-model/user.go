package query

import (
	"time"

	"github.com/google/uuid"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	permissionQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"gorm.io/gorm"
)

type User struct {
	CreatedAt           time.Time                                      `gorm:"column:created_at;"`
	UpdatedAt           time.Time                                      `gorm:"column:updated_at;"`
	LastAccessAt        time.Time                                      `gorm:"column:last_access_at"`
	SSN                 *string                                        `gorm:"column:ssn;"`
	DeletedAt           *gorm.DeletedAt                                `gorm:"column:deleted_at"`
	NameTH              string                                         `gorm:"column:name_th;"`
	NameEN              string                                         `gorm:"column:name_en;"`
	Email               string                                         `gorm:"column:email;"`
	Type                string                                         `gorm:"column:type;"`
	Status              string                                         `gorm:"column:status;"`
	CurrentRoleID       uint                                           `gorm:"column:current_role_id"`
	SystemPermission    permissionQuery.PermissionSystemQueryEntity    `gorm:"foreignKey:SystemPermissionUID;references:UID"`
	CoursePermission    []permissionQuery.PermissionCourseQueryEntity  `gorm:"foreignKey:UserUID;references:UID"`
	ProgramPermission   []permissionQuery.PermissionProgramQueryEntity `gorm:"foreignKey:UserUID;references:UID"`
	UID                 uuid.UUID                                      `gorm:"column:uid;"`
	SystemPermissionUID uuid.UUID                                      `gorm:"column:system_permission_uid;"`
	FacultyID           uint                                           `gorm:"column:faculty_id;"`
}

type UserQueryEntity struct {
	CreatedAt               *time.Time                                     `gorm:"column:created_at;"`
	UpdatedAt               *time.Time                                     `gorm:"column:updated_at;"`
	LastAccessAt            *time.Time                                     `gorm:"column:last_access_at"`
	SSN                     *string                                        `gorm:"column:ssn;"`
	DeletedAt               *gorm.DeletedAt                                `gorm:"column:deleted_at"`
	NameTH                  *string                                        `gorm:"column:name_th;"`
	NameEN                  *string                                        `gorm:"column:name_en;"`
	Email                   *string                                        `gorm:"column:email;"`
	Type                    *string                                        `gorm:"column:type;"`
	Status                  *string                                        `gorm:"column:status;"`
	SystemPermission        *permissionQuery.PermissionSystemQueryEntity   `gorm:"foreignKey:SystemPermissionUID;references:UID"`
	CoursePermission        []permissionQuery.PermissionCourseQueryEntity  `gorm:"foreignKey:UserUID;references:UID"`
	ProgramPermission       []permissionQuery.PermissionProgramQueryEntity `gorm:"foreignKey:UserUID;references:UID"`
	UID                     *uuid.UUID                                     `gorm:"column:uid;"`
	SystemPermissionUID     *uuid.UUID                                     `gorm:"column:system_permission_uid;"`
	FacultyID               *uint                                          `gorm:"column:faculty_id;"`
	Faculty                 *query.Faculty                                 `gorm:"foreignKey:ID;references:FacultyID"`
	CurrentRoleID           *uint                                          `gorm:"column:current_role_id"`
	ProgramApprovalMinLevel uint                                           `gorm:"column:program_approval_min_level;type:smallint;default:0"`
	ProgramApprovalMaxLevel uint                                           `gorm:"column:program_approval_max_level;type:smallint;default:0"`
}

func (e *UserQueryEntity) TableName() string {
	return "users"
}
