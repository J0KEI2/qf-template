package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type GetPossibleRolesResponseDto struct {
	Items []GetRoleResponseDto `json:"items"`
}

type GetRoleResponseDto struct {
	ID                   *uint   `json:"id"`
	RoleNameTH           *string `json:"role_name_th"`
	RoleNameEN           *string `json:"role_name_en"`
	ProgramRoleType      *uint   `json:"program_role_type"`
	CourseRoleType       *uint   `json:"course_role_type"`
	SystemLevel          *uint   `json:"system_level"`
	ProgramApprovalLevel *uint   `json:"program_approval_level"`
	CourseApprovalLevel  *uint   `json:"course_approval_level"`
	ProgramActionLevel   *uint   `json:"program_action_level"`
	CourseActionLevel    *uint   `json:"course_action_level"`
	ProgramAccessLevel   *uint   `json:"program_access_level"`
	CourseAccessLevel    *uint   `json:"course_access_level"`
}

type CreateOrUpdateRoleRequestDto struct {
	ID                   *uint   `json:"id"`
	RoleNameTH           *string `json:"role_name_th"`
	RoleNameEN           *string `json:"role_name_en"`
	ProgramRoleType      *uint   `json:"program_role_type"`
	CourseRoleType       *uint   `json:"course_role_type"`
	SystemLevel          *uint   `json:"system_level"`
	ProgramApprovalLevel *uint   `json:"program_approval_level"`
	CourseApprovalLevel  *uint   `json:"course_approval_level"`
	ProgramActionLevel   *uint   `json:"program_action_level"`
	CourseActionLevel    *uint   `json:"course_action_level"`
}

type CreateSettingUserCoreRoleRequestDto struct {
	UserID *uuid.UUID `json:"user_id"`
	RoleID *uint      `json:"role_id"`
}

type GetSettingUserRoleResponseDto struct {
	Items []SettingUserRoleDto `json:"items"`
	*models.PaginationOptions
}

type SettingUserRoleDto struct {
	UserID        *uuid.UUID `json:"user_id"`
	NameTH        *string    `json:"name_th"`
	NameEN        *string    `json:"name_en"`
	FacultyNameTH *string    `json:"faculty_name_th"`
	FacultyNameEN *string    `json:"faculty_name_en"`
	Email         *string    `json:"email"`
	CreatedAt     *time.Time `json:"created_at"`
}

type CreateSettingUserFacultyRoleRequestDto struct {
	UserID    *uuid.UUID `json:"user_id"`
	RoleID    *uint      `json:"role_id"`
	FacultyID *uint      `json:"faculty_id"`
}
