package models

import (
	"time"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"gorm.io/gorm"
)

type UserFetchQuery struct {
	UID                 *uuid.UUID `gorm:"column:uid;type:uuid;primaryKey" json:"uid"`
	Email               *string    `gorm:"column:email;type:varchar;size:255;not null;uniqueIndex" json:"email"`
	SSN                 *string    `gorm:"column:ssn;type:varchar;size:255" json:"ssn,omitempty"`
	SystemPermissionUID *uuid.UUID `gorm:"column:system_permission_uid;type:uuid;not null" json:"system_permission_uid"`
	FacultyID           *uint      `gorm:"column:faculty_id;type:int;not null" json:"faculty_id"`
	Type                *string    `gorm:"column:type;type:varchar;size:255;not null" json:"type"`
	Status              *string    `gorm:"column:status;type:varchar;size:255;not null" json:"status"`
	OrderBy             *string
	Direction           *string
	Limit               *int
	Offset              *int
}

type UserFetchModel struct {
	CreatedAt               time.Time  `gorm:"column:created_at;type:timestamp;CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt               time.Time  `gorm:"column:updated_at;type:timestamp;CURRENT_TIMESTAMP()" json:"updated_at"`
	LastAccessAt            time.Time  `gorm:"column:last_access_at;type:UserStatus;not null" json:"last_access_at"`
	Email                   *string    `gorm:"column:email;type:varchar;size:255;not null;uniqueIndex" json:"email"`
	TitleTH                 *string    `gorm:"column:title_th;type:varchar;size:255" json:"title_th"`
	TitleEN                 *string    `gorm:"column:title_en;type:varchar;size:255" json:"title_en"`
	FirstNameTH             *string    `gorm:"column:first_name_th;type:varchar;size:255" json:"first_name_th"`
	FirstNameEN             *string    `gorm:"column:first_name_en;type:varchar;size:255" json:"first_name_en"`
	LastNameTH              *string    `gorm:"column:last_name_th;type:varchar;size:255" json:"last_name_th"`
	LastNameEN              *string    `gorm:"column:last_name_en;type:varchar;size:255" json:"last_name_en"`
	SSN                     *string    `gorm:"column:ssn;type:varchar;size:255" json:"ssn,omitempty"`
	Type                    *string    `gorm:"column:type;type:varchar;size:255;not null" json:"type"`
	Status                  *string    `gorm:"column:status;type:varchar;size:255;not null" json:"status"`
	UID                     *uuid.UUID `gorm:"column:uid;type:uuid;primaryKey" json:"uid"`
	SystemPermissionUID     *uuid.UUID `gorm:"column:system_permission_uid;type:uuid;not null" json:"system_permission_uid"`
	FacultyID               *uint      `gorm:"column:faculty_id;type:int" json:"faculty_id"`
	FacultyName             *string    `gorm:"column:faculty_name;type:varchar;size:255" json:"faculty_name"`
	SystemPermissionName    *string    `gorm:"column:system_permission_name;type:varchar;size:255" json:"system_permission_name"`
	ProgramApprovalMinLevel uint       `gorm:"column:program_approval_min_level;type:smallint;default:0" json:"program_approval_min_level"`
	ProgramApprovalMaxLevel uint       `gorm:"column:program_approval_max_level;type:smallint;default:0" json:"program_approval_max_level"`
	CurrentRoleID           uint       `gorm:"column:current_role_id;type:uint" json:"current_role_id"`
}
type UserFetchWithSystemRoleModel struct {
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	LastAccessAt        time.Time `json:"last_access_at"`
	Email               string    `json:"email"`
	SSN                 string    `json:"ssn"`
	Type                string    `json:"type"`
	UserStatus          string    `json:"user_status"`
	UID                 uuid.UUID `json:"uid"`
	SystemPermissionUID uuid.UUID `json:"system_permission_uid"`
	FacultyID           uint      `json:"faculty_id"`
	RoleNameTH          string    `json:"role_name_th"`
	RoleNameEN          string    `json:"role_name_en"`
	ViewMyProfile       bool      `json:"view_my_profile"`
	UpdateMyProfile     bool      `json:"update_my_profile"`
	ListAllLecture      bool      `json:"list_all_lecture"`
	CreateLecture       bool      `json:"create_lecture"`
	UpdateLecture       bool      `json:"update_lecture"`
	DeleteLecture       bool      `json:"delete_lecture"`
	Status              bool      `json:"status"`
}
type UserCreateQuery struct {
	SSN                 *string   `json:"ssn,omitempty"`
	LastnameTh          string    `json:"lastname_th"`
	FirstnameEn         string    `json:"firstname_en"`
	LastnameEn          string    `json:"lastname_en"`
	TitleEn             string    `json:"title_en"`
	Email               string    `json:"email"`
	FirstnameTh         string    `json:"firstname_th"`
	TitleTh             string    `json:"title_th"`
	Type                string    `json:"type"`
	SystemPermissionUID uuid.UUID `json:"system_permission_uid"`
	FacultyID           uint      `json:"faculty_id"`
}

type UserFetchWithRelationQueryModel struct {
	UID                 *uuid.UUID
	Name                *string
	Email               *string
	FacultyNameTH       *string
	RoleNameEN          *string
	SSN                 *string
	SystemPermissionUID *uuid.UUID
	FacultyID           *uuid.UUID
	Type                *string
	Status              *string
	OrderBy             *string
	Direction           *string
	Limit               *int
	Offset              *int
}

type UserFetchWithRelationModel struct {
	CreatedAt           time.Time `gorm:"column:created_at;type:timestamp;CURRENT_TIMESTAMP()" json:"created_at"`
	LastAccessAt        time.Time `gorm:"column:last_access_at;type:UserStatus;not null" json:"lastAccessAt"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:timestamp;CURRENT_TIMESTAMP()" json:"updated_at"`
	NameTH              string    `gorm:"column:name_th;type:varchar;size:255;not null" json:"name_th"`
	NameEN              string    `gorm:"column:name_en;type:varchar;size:255;not null" json:"name_en"`
	Status              string    `gorm:"column:status;type:varchar;size:255;not null" json:"status"`
	Type                string    `gorm:"column:type;type:varchar;size:255;not null" json:"type"`
	SSN                 string    `gorm:"column:ssn;type:varchar;size:255" json:"ssn,omitempty"`
	RoleNameTH          string    `gorm:"column:role_name_th;type:varchar;size:200;not null" json:"roleNameTH"`
	RoleNameEN          string    `gorm:"column:role_name_en;type:varchar;size:200;not null" json:"roleNameEN"`
	FacultyNameTH       string    `gorm:"column:faculty_name_th;type:varchar;size:255;not null" json:"facultyNameTH,omitempty"`
	FacultyNameEN       string    `gorm:"column:faculty_name_en;type:varchar;size:255;not null" json:"facultyNameEN,omitempty"`
	Email               string    `gorm:"column:email;type:varchar;size:255;not null;uniqueIndex" json:"email"`
	SystemPermissionUID uuid.UUID `gorm:"column:system_permission_uid;type:uuid;not null" json:"system_permission_uid"`
	UID                 uuid.UUID `gorm:"column:uid;type:uuid;primaryKey" json:"uid"`
	FacultyID           uint      `gorm:"column:faculty_id;type:int" json:"faculty_id"`
}

type PatchUserRequest struct {
	TitleTH             string    `json:"title_th"`
	TitleEN             string    `json:"title_en"`
	FirstNameTH         string    `json:"first_name_th"`
	FirstNameEN         string    `json:"first_name_en"`
	LastNameTH          string    `json:"last_name_th"`
	LastNameEN          string    `json:"last_name_en"`
	SSN                 string    `json:"ssn,omitempty"`
	Type                string    `json:"type"`
	Status              string    `json:"status"`
	SystemPermissionUID uuid.UUID `json:"system_permission_uid"`
	FacultyID           uint      `json:"faculty_id"`
}

type User struct {
	ID        string         `json:"id" gorm:"size:32;primaryKey"`
	Password  string         `json:"password" gorm:"size:64"`
	FullName  string         `json:"full_name" gorm:"size:127;index"`
	Address   string         `json:"address" gorm:"type:text"`
	CreatedAt string         `json:"created_at,omitempty" gorm:"autoCreateTime;index"`
	UpdatedAt string         `json:"updated_at,omitempty" gorm:"autoUpdateTime;index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
type UserResponse struct {
	helpers.ResponseForm
}

type QueryUpdateUserQueryBuilder struct {
	UID *uuid.UUID
}

func (m QueryUpdateUserQueryBuilder) GetDBQuery(db *gorm.DB) *gorm.DB {
	if m.UID != nil {
		db = db.Where("uid = ?", *m.UID)
	}
	return db
}
