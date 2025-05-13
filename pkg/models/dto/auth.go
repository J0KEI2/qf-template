package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserID string     `json:"user_id"`
	Role   RoleClaims `json:"role"`
}

type RoleClaims struct {
	ID                   uint    `json:"id"`
	RoleNameTH           *string `json:"role_name_th"`
	RoleNameEN           *string `json:"role_name_en"`
	ProgramRoleType      *uint   `json:"program_role_type"`
	CourseRoleType       *uint   `json:"course_role_type"`
	ProgramApprovalLevel *uint   `json:"program_approval_level"`
	CourseApprovalLevel  *uint   `json:"course_approval_level"`
	ProgramActionLevel   *uint   `json:"program_action_level"`
	CourseActionLevel    *uint   `json:"course_action_level"`
}

type SwitchRoleRequest struct {
	RoleID uint `json:"role_id" query:"role_id"`
}
