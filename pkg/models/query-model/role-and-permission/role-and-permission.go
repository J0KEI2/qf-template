package query

import (
	"time"

	"github.com/google/uuid"
	userQuery "github.com/zercle/kku-qf-services/pkg/models/query-model"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

type MapCoreRolesQueryEntity struct {
	ID        *uint                      `gorm:"column:id"`
	UserID    *uuid.UUID                 `gorm:"column:user_id"`
	User      *userQuery.UserQueryEntity `gorm:"foreignKey:UID;references:UserID"`
	RoleID    *uint                      `gorm:"column:role_id"`
	Role      *RoleQueryEntity           `gorm:"foreignKey:ID;references:RoleID"`
	CreatedAt *time.Time                 `gorm:"column:created_at"`
	UpdatedAt *time.Time                 `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt            `gorm:"column:deleted_at"`
}

func (MapCoreRolesQueryEntity) TableName() string {
	return "map_core_roles"
}

type MapFacultiesRolesQueryEntity struct {
	ID        *uint                      `gorm:"column:id"`
	UserID    *uuid.UUID                 `gorm:"column:user_id"`
	User      *userQuery.UserQueryEntity `gorm:"foreignKey:UID;references:UserID"`
	RoleID    *uint                      `gorm:"column:role_id"`
	Role      *RoleQueryEntity           `gorm:"foreignKey:ID;references:RoleID"`
	FacultyID *uint                      `gorm:"column:faculty_id"`
	Faculty   *query.Faculty             `gorm:"foreignKey:ID;references:FacultyID"`
	CreatedAt *time.Time                 `gorm:"column:created_at"`
	UpdatedAt *time.Time                 `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt            `gorm:"column:deleted_at"`
}

func (MapFacultiesRolesQueryEntity) TableName() string {
	return "map_faculties_roles"
}

type MapProgramsRolesQueryEntity struct {
	ID        *uint                                `gorm:"column:id"`
	UserID    *uuid.UUID                           `gorm:"column:user_id"`
	RoleID    *uint                                `gorm:"column:role_id"`
	Role      *RoleQueryEntity                     `gorm:"foreignKey:ID;references:RoleID"`
	User      *userQuery.UserQueryEntity           `gorm:"foreignKey:UID;references:UserID"`
	ProgramID *uuid.UUID                           `gorm:"column:program_id"`
	Program   *programQuery.ProgramMainQueryEntity `gorm:"foreignKey:ID;references:ProgramID"`
	CreatedAt *time.Time                           `gorm:"column:created_at"`
	UpdatedAt *time.Time                           `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt                      `gorm:"column:deleted_at"`
}

func (MapProgramsRolesQueryEntity) TableName() string {
	return "map_programs_roles"
}

type RoleQueryEntity struct {
	ID                   *uint                          `gorm:"column:id"`
	RoleNameTH           *string                        `gorm:"column:role_name_th"`
	RoleNameEN           *string                        `gorm:"column:role_name_en"`
	ProgramRoleType      *uint                          `gorm:"column:program_role_type"`
	CourseRoleType       *uint                          `gorm:"column:course_role_type"`
	SystemLevel          *uint                          `gorm:"column:system_level"`
	ProgramApprovalLevel *uint                          `gorm:"column:program_approval_level"`
	CourseApprovalLevel  *uint                          `gorm:"column:course_approval_level"`
	ProgramActionLevel   *uint                          `gorm:"column:program_action_level"`
	CourseActionLevel    *uint                          `gorm:"column:course_action_level"`
	ProgramAccessLevel   *uint                          `gorm:"column:program_access_level"`
	CourseAccessLevel    *uint                          `gorm:"column:course_access_level"`
	MapCoreRoles         []MapCoreRolesQueryEntity      `gorm:"foreignKey:RoleID;references:ID"`
	MapFacultiesRoles    []MapFacultiesRolesQueryEntity `gorm:"foreignKey:RoleID;references:ID"`
	MapProgramRoles      []MapProgramsRolesQueryEntity  `gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt            *time.Time                     `gorm:"column:created_at"`
	UpdatedAt            *time.Time                     `gorm:"column:updated_at"`
	DeletedAt            *gorm.DeletedAt                `gorm:"column:deleted_at"`
}

func (RoleQueryEntity) TableName() string {
	return "roles"
}
