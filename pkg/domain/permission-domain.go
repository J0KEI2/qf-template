package domain

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

type PermissionUseCase interface {
	CreateNewSystemPermission(request dto.CreatePermissionSystemDto) (result *dto.PermissionSystemResponseDto, err error)
	DeleteSystemPermission(uid uuid.UUID) (err error)
	GetOneSystemPermission(uid uuid.UUID) (result *dto.PermissionSystemResponseDto, err error)
	GetAllSystemPermission() (result []dto.PermissionSystemResponseDto, err error)
	UpdateSystemPermission(uid uuid.UUID, update dto.UpdatePermissionSystemDto) (result *dto.PermissionSystemResponseDto, err error)
	GetAllProgramPermissionByUser(uid uuid.UUID) (result []dto.PermissionProgramResponseDto, err error)
	GetAllProgramPermissionByProgram(uid uuid.UUID) (result []dto.PermissionProgramResponseDto, err error)
	CreateOrUpdateProgramPermission(permissionPrograms []dto.UpdatePermissionProgramDto) (result []dto.UpdatePermissionProgramDto, err error)
	DeleteProgramPermission(deleteIdList []int) (err error)
	GetCoreRolesByUid(uid uuid.UUID, options models.PaginationOptions) (result *dto.GetMapCoreRolesResponseDto, err error)
	CreateCoreRoles(request dto.CreateMapCoreRolesRequestDto) (result *dto.MapCoreRoles, err error)
	DeleteCoreRoles(mapId uint) (err error)
	GetFacultyRolesByUid(uid uuid.UUID, options models.PaginationOptions) (result *dto.GetMapFacultyRolesResponseDto, err error)
	CreateFacultyRoles(request dto.CreateMapFacultyRolesRequestDto) (result *dto.MapFacultyRoles, err error)
	DeleteFacultyRoles(mapId uint) (err error)
	CreateProgramRoles(request dto.CreateMapProgramRolesRequestDto) (result *dto.MapProgramRoles, err error)
	DeleteProgramRoles(mapId uint) (err error)
	GetProgramRolesByUid(userUID *uuid.UUID, programUID *uuid.UUID, options models.PaginationOptions) (result *dto.GetMapProgramRolesResponseDto, err error)
	GetProgramRolesByProgramUid(programUID *uuid.UUID, programType *string, options models.PaginationOptions) (result *dto.GetMapUserProgramRolesResponseDto, err error)
}

type PermissionRepository interface {
	DbPermissionSVCMigrator() (err error)
}
