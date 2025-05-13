package dto

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type MapCoreRoles struct {
	ID         *uint      `json:"id"`
	UserID     *uuid.UUID `json:"user_id"`
	RoleID     *uint      `json:"role_id"`
	RoleNameTH *string    `json:"role_name_th"`
	RoleNameEN *string    `json:"role_name_en"`
}

type GetMapCoreRolesResponseDto struct {
	Items []MapCoreRoles `json:"items"`
	*models.PaginationOptions
}

type CreateMapCoreRolesRequestDto struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID uint      `json:"role_id"`
}

type MapFacultyRoles struct {
	ID            *uint      `json:"id"`
	UserID        *uuid.UUID `json:"user_id"`
	FacultyID     *uint      `json:"faculty_id"`
	FacultyNameTH *string    `json:"faculty_name_th"`
	FacultyNameEN *string    `json:"faculty_name_en"`
	RoleID        *uint      `json:"role_id"`
	RoleNameTH    *string    `json:"role_name_th"`
	RoleNameEN    *string    `json:"role_name_en"`
}

type GetMapFacultyRolesResponseDto struct {
	Items []MapFacultyRoles `json:"items"`
	*models.PaginationOptions
}

type CreateMapFacultyRolesRequestDto struct {
	UserID    uuid.UUID `json:"user_id"`
	FacultyID uint      `json:"faculty_id"`
	RoleID    uint      `json:"role_id"`
}

type MapProgramRoles struct {
	ID            *uint      `json:"id"`
	UserID        *uuid.UUID `json:"user_id"`
	ProgramID     *uuid.UUID `json:"program_id"`
	ProgramNameTH *string    `json:"program_name_th"`
	ProgramNameEN *string    `json:"program_name_en"`
	RoleID        *uint      `json:"role_id"`
	RoleNameTH    *string    `json:"role_name_th"`
	RoleNameEN    *string    `json:"role_name_en"`
}

type MapUserProgramRoles struct {
	ID            *uint      `json:"id"`
	UserID        *uuid.UUID `json:"user_id"`
	ProgramID     *uuid.UUID `json:"program_id"`
	RoleID        *uint      `json:"role_id"`
	RoleNameTH    *string    `json:"role_name_th"`
	RoleNameEN    *string    `json:"role_name_en"`
	NameTH        *string    `json:"name_th"`
	NameEN        *string    `json:"name_en"`
	Email         *string    `json:"email"`
	FacultyNameEN *string    `json:"faculty_name_en"`
	FacultyNameTH *string    `json:"faculty_name_th"`
}

type GetMapProgramRolesResponseDto struct {
	Items []MapProgramRoles `json:"items"`
	*models.PaginationOptions
}

type GetMapUserProgramRolesResponseDto struct {
	Items []MapUserProgramRoles `json:"items"`
	*models.PaginationOptions
}

type CreateMapProgramRolesRequestDto struct {
	UserID    uuid.UUID `json:"user_id"`
	ProgramID uuid.UUID `json:"program_id"`
	RoleID    uint      `json:"role_id"`
}
