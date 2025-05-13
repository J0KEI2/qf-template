package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u *PermissionUsecase) GetProgramRolesByUid(userUID *uuid.UUID, programUID *uuid.UUID, options models.PaginationOptions) (result *dto.GetMapProgramRolesResponseDto, err error) {
	mapProgramRoleStatement := query.MapProgramsRolesQueryEntity{
		UserID:    userUID,
		ProgramID: programUID,
	}

	mapProgramRoleQuery := make([]query.MapProgramsRolesQueryEntity, 0)

	err = u.CommonRepository.GetList(mapProgramRoleStatement, &mapProgramRoleQuery, &options, "Role", "Program.ProgramGeneralDetail")

	if err != nil {
		return nil, err
	}

	mapProgramRoleResult := make([]dto.MapProgramRoles, 0)

	for _, mapProgramRole := range mapProgramRoleQuery {
		programRole := dto.MapProgramRoles{
			ID:         mapProgramRole.ID,
			UserID:     mapProgramRole.UserID,
			RoleID:     mapProgramRole.RoleID,
			RoleNameTH: mapProgramRole.Role.RoleNameTH,
			RoleNameEN: mapProgramRole.Role.RoleNameEN,
			ProgramID:  mapProgramRole.ProgramID,
		}
		if mapProgramRole.Program != nil && mapProgramRole.Program.ProgramGeneralDetail != nil {
			programRole.ProgramNameTH = mapProgramRole.Program.ProgramGeneralDetail.ProgramNameTH
			programRole.ProgramNameEN = mapProgramRole.Program.ProgramGeneralDetail.ProgramNameEN
		}
		mapProgramRoleResult = append(mapProgramRoleResult, programRole)
	}

	result = &dto.GetMapProgramRolesResponseDto{
		Items:             mapProgramRoleResult,
		PaginationOptions: &options,
	}
	return result, nil
}
