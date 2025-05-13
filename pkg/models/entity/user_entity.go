package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"gorm.io/gorm"
)

const (
	EmailName     = "email"
	FirstNameTH   = "firstname_th"
	FirstNameEN   = "firstname_en"
	FacultyNameTH = "faculty_name_th"
	RoleNameEN    = "role_name_en"
)

type UserFetchQueryEntity struct {
	UID              *uuid.UUID
	Email            *string
	Name             *string
	RoleSystemNameEN *string
	FacultyNameTH    *string
	Status           *string
	OrderBy          *string
	Direction        *string
	Limit            *int
	Page             *int
}

type UserQueryEntityWithRole struct {
	CreatedAt               *time.Time      `gorm:"column:created_at;"`
	UpdatedAt               *time.Time      `gorm:"column:updated_at;"`
	LastAccessAt            *time.Time      `gorm:"column:last_access_at"`
	SSN                     *string         `gorm:"column:ssn;"`
	DeletedAt               *gorm.DeletedAt `gorm:"column:deleted_at"`
	NameTH                  *string         `gorm:"column:name_th;"`
	NameEN                  *string         `gorm:"column:name_en;"`
	Email                   *string         `gorm:"column:email;"`
	Type                    *string         `gorm:"column:type;"`
	Status                  *string         `gorm:"column:status;"`
	UID                     *uuid.UUID      `gorm:"column:uid;"`
	SystemPermissionUID     *uuid.UUID      `gorm:"column:system_permission_uid;"`
	FacultyID               *uint           `gorm:"column:faculty_id;"`
	CurrentRoleID           *uint           `gorm:"column:current_role_id"`
	ProgramApprovalMinLevel *uint           `gorm:"column:program_approval_min_level;type:smallint;default:0"`
	ProgramApprovalMaxLevel *uint           `gorm:"column:program_approval_max_level;type:smallint;default:0"`
	RoleNameTH              *string         `gorm:"column:role_name_th"`
	RoleNameEN              *string         `gorm:"column:role_name_en"`
}

type UserListFetchEntity struct {
	Pagination *models.PaginationOptions
	Data       []UserFetchEntity
}

type UserFetchListEntity struct {
	Options *models.PaginationOptions `json:"options"`
	Items   []UserFetchEntity         `json:"items"`
}

type UserFetchEntity struct {
	CreatedAt           time.Time `json:"created_at"`
	LastAccessAt        time.Time `json:"last_access_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Email               string    `json:"email"`
	FacultyNameTH       string    `json:"faculty_name_th"`
	FirstnameTh         string    `json:"firstname_th"`
	TitleTh             string    `json:"title_th"`
	Status              string    `json:"status"`
	SSN                 string    `json:"ssn"`
	HRKey               string    `json:"hr_key"`
	RoleNameTH          string    `json:"role_name_th"`
	RoleNameEN          string    `json:"role_name_en"`
	FacultyNameEN       string    `json:"faculty_name_en"`
	Type                string    `json:"type"`
	LastnameTH          string    `json:"lastname_th"`
	TitleEn             string    `json:"title_en"`
	FirstnameEn         string    `json:"firstname_en"`
	LastnameEn          string    `json:"lastname_en"`
	NameTH              string    `json:"name_th"`
	NameEN              string    `json:"name_en"`
	FacultyID           uint      `json:"faculty_id"`
	SystemPermissionUID uuid.UUID `json:"system_permission_uid"`
	UID                 uuid.UUID `json:"uid"`
	CurrentRoleID       uint      `json:"current_role_id"`
}

type QueryUpdateUser struct {
	ID *uuid.UUID
}
