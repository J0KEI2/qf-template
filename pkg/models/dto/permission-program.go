package dto

import "github.com/google/uuid"

type PermissionProgramResponseDto struct {
	ID            *int       `json:"id"`
	UserUID       *uuid.UUID `json:"user_uid"`
	ProgramUID    *uuid.UUID `json:"program_uid"`
	FacultyName   *string    `json:"faculty"`
	ProgramName   *string    `json:"program_name"`
	BranchName    *string    `json:"branch_name"`
	ProgramType   *string    `json:"program_type"`
	Name          *string    `json:"name"`
	Email         *string    `json:"email"`
	Accessibility *uint      `json:"accessibility"`
}

type CreatePermissionProgramDto struct {
	RoleNameTH                string   `json:"role_name_th"`
	RoleNameEN                string   `json:"role_name_en"`
	PageAccessibility         []string `json:"page_accessibility"`
	ProgramAccessibility      []string `json:"program_accessibility"`
	CourseAccessibility       []string `json:"course_accessibility"`
	ProgramAccessibilityLevel uint     `json:"program_accessibility_level"`
	CourseAccessibilityLevel  uint     `json:"course_accessibility_level"`
	UAMControl                bool     `json:"uam_control"`
	CanComment                bool     `json:"can_comment"`
	CanApproved               bool     `json:"can_approve"`
}

type UpdatePermissionProgramDto struct {
	ID            *int       `json:"id"`
	UserUID       *uuid.UUID `json:"user_uid"`
	ProgramUID    *uuid.UUID `json:"program_uid"`
	Accessibility *uint      `json:"accessibility"`
}
