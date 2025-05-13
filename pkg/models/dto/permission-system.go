package dto

type PermissionSystemResponseDto struct {
	UID                       string   `json:"uid"`
	RoleNameTH                *string  `json:"role_name_th"`
	RoleNameEN                *string  `json:"role_name_en"`
	PageAccessibility         []string `json:"page_accessibility"`
	ProgramAccessibility      []string `json:"program_accessibility"`
	CourseAccessibility       []string `json:"course_accessibility"`
	ProgramAccessibilityLevel *uint    `json:"program_accessibility_level"`
	CourseAccessibilityLevel  *uint    `json:"course_accessibility_level"`
	UAMControl                bool     `json:"uam_control"`
	CanComment                bool     `json:"can_comment"`
	CanApprove                bool     `json:"can_approve"`
	ProgramApprovalMinLevel   uint     `json:"program_approval_min_level"`
	ProgramApprovalMaxLevel   uint     `json:"program_approval_max_level"`
}

type CreatePermissionSystemDto struct {
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

type UpdatePermissionSystemDto struct {
	RoleNameTH                *string  `json:"role_name_th"`
	RoleNameEN                *string  `json:"role_name_en"`
	PageAccessibility         []string `json:"page_accessibility"`
	ProgramAccessibility      []string `json:"program_accessibility"`
	CourseAccessibility       []string `json:"course_accessibility"`
	ProgramAccessibilityLevel *uint    `json:"program_accessibility_level"`
	CourseAccessibilityLevel  *uint    `json:"course_accessibility_level"`
	UAMControl                bool     `json:"uam_control"`
	CanComment                bool     `json:"can_comment"`
	CanApproved               bool     `json:"can_approve"`
}
