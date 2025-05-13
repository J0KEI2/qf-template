package models

import (
	"github.com/google/uuid"
)

type UpdateUserRoleRequest struct {
	EducationBackGround     *string     `json:"education_background"`
	Position                []string    `json:"position"`
	UserUID                 uuid.UUID   `json:"user_uid"`
	YearOfAcceptingPosition int         `json:"year_of_accepting_position"`
	SystemPermissionUID     uuid.UUID   `json:"system_permission_uid"`
	UserRoleUID             uuid.UUID   `json:"user_role_uid"`
	RoleCourseUID           []uuid.UUID `json:"role_course_uid"`
	RoleProgramUID          []uuid.UUID `json:"role_program_uid"`
}

type UserDetailPostion struct {
	Position string `json:"position"`
	Year     int    `json:"year"`
}
