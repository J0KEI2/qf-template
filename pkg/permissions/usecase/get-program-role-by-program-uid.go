package usecase

import (
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/permission"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u *PermissionUsecase) GetProgramRolesByProgramUid(programUID *uuid.UUID, programType *string, options models.PaginationOptions) (result *dto.GetMapUserProgramRolesResponseDto, err error) {
	var roleId *uint = nil
	switch strings.ToLower(pointer.GetString(programType)) {
	case constant.ROLE_TYPE_OWNER:
		{
			roleQuery := query.RoleQueryEntity{
				RoleNameEN: pointer.ToString(constant.ROLE_PROGRAM_OWNER),
			}

			u.CommonRepository.GetFirst(&roleQuery)
			roleId = roleQuery.ID
		}
	case constant.ROLE_TYPE_LECTURER:
		{
			roleQuery := query.RoleQueryEntity{
				RoleNameEN: pointer.ToString(constant.ROLE_PROGRAM_LECTURER),
			}

			u.CommonRepository.GetFirst(&roleQuery)
			roleId = roleQuery.ID
		}
	default:
		{
			roleId = nil
		}
	}

	mapProgramRoleStatement := query.MapProgramsRolesQueryEntity{
		ProgramID: programUID,
		RoleID:    roleId,
	}

	mapProgramRoleQuery := make([]query.MapProgramsRolesQueryEntity, 0)

	options.SetSearchFields([]string{`"User".email`, `"User".name_th`, `"User".name_en`})

	err = u.CommonRepository.GetList(mapProgramRoleStatement, &mapProgramRoleQuery, &options, "Role", "User", "User.Faculty")

	if err != nil {
		return nil, err
	}

	mapProgramRoleResult := make([]dto.MapUserProgramRoles, 0)

	for _, mapProgramRole := range mapProgramRoleQuery {
		programRole := dto.MapUserProgramRoles{
			ID:        mapProgramRole.ID,
			UserID:    mapProgramRole.UserID,
			RoleID:    mapProgramRole.RoleID,
			ProgramID: mapProgramRole.ProgramID,
		}
		if mapProgramRole.Role != nil {
			programRole.RoleNameTH = mapProgramRole.Role.RoleNameTH
			programRole.RoleNameEN = mapProgramRole.Role.RoleNameEN
		}
		if mapProgramRole.User != nil {
			programRole.NameEN = mapProgramRole.User.NameEN
			programRole.NameTH = mapProgramRole.User.NameTH
			programRole.Email = mapProgramRole.User.Email
			if mapProgramRole.User.Faculty != nil {
				programRole.FacultyNameEN = mapProgramRole.User.Faculty.FacultyNameEN
				programRole.FacultyNameTH = mapProgramRole.User.Faculty.FacultyNameTH
			}
		}
		mapProgramRoleResult = append(mapProgramRoleResult, programRole)
	}

	result = &dto.GetMapUserProgramRolesResponseDto{
		Items:             mapProgramRoleResult,
		PaginationOptions: &options,
	}
	return result, nil
}
