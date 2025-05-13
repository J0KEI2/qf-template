package domain

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

type RoleUseCase interface {
	GetPossibleRole(userID uuid.UUID) (result *dto.GetPossibleRolesResponseDto, err error)
	GetIfUserHasRoles(userID uuid.UUID) (ok bool, err error)
	CheckIfUserHasRole(userID uuid.UUID, roleID *uint) (ok bool, err error)
	GetRoleByID(roleID uint) (result *dto.GetRoleResponseDto, err error)
	UpdateUserCurrentRole(userID uuid.UUID, roleID uint) (err error)
	CreateOrUpdateRole(request dto.CreateOrUpdateRoleRequestDto) (err error)
	DeleteRoleByID(roleID uint) (err error)
	CreateSettingUserCoreRole(request dto.CreateSettingUserCoreRoleRequestDto) (err error)
	CreateSettingUserFacultyRole(request dto.CreateSettingUserFacultyRoleRequestDto) (err error)

	DeleteSettingUserFacultyRole(userID *uuid.UUID, roleID, facultyID *uint) (err error)
	DeleteSettingUserCoreRole(userID *uuid.UUID, roleID *uint) (err error)

	GetSettingUserCoreRole(roleID *uint, options models.PaginationOptions) (result *dto.GetSettingUserRoleResponseDto, err error)
	GetSettingUserFacultyRole(roleID, facultyID *uint, options models.PaginationOptions) (result *dto.GetSettingUserRoleResponseDto, err error)
}

type RoleRepository interface {
	DbRoleSVCMigrator() (err error)
	DeleteSettingUserCoreRole(userID uuid.UUID, roleID uint) (err error)
	DeleteSettingUserFacultyRole(userID uuid.UUID, roleID, facultyID uint) (err error)
}
