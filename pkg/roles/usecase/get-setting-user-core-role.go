package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	queryRoleAndPermission "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u roleUsecase) GetSettingUserCoreRole(roleID *uint, options models.PaginationOptions) (result *dto.GetSettingUserRoleResponseDto, err error) {
	if *roleID != 0 {
		roleResponse := []dto.SettingUserRoleDto{}
		mapCoreRoleList := []queryRoleAndPermission.MapCoreRolesQueryEntity{}
		mapCoreRoleQuery := queryRoleAndPermission.MapCoreRolesQueryEntity{
			RoleID: roleID,
		}

		options.SetSearchFields([]string{`"User".email`, `"User".name_th`, `"User".name_en`})

		if err = u.CommonRepository.GetList(&mapCoreRoleQuery, &mapCoreRoleList, &options, "Role", "User", "User.Faculty"); err != nil {
			return nil, err
		}

		for _, mapCoreRoleItem := range mapCoreRoleList {
			roleResponse = append(roleResponse, dto.SettingUserRoleDto{
				UserID:        mapCoreRoleItem.UserID,
				NameTH:        mapCoreRoleItem.User.NameTH,
				NameEN:        mapCoreRoleItem.User.NameEN,
				FacultyNameTH: mapCoreRoleItem.User.Faculty.FacultyNameTH,
				FacultyNameEN: mapCoreRoleItem.User.Faculty.FacultyNameEN,
				Email:         mapCoreRoleItem.User.Email,
				CreatedAt:     mapCoreRoleItem.CreatedAt,
			})
		}

		response := &dto.GetSettingUserRoleResponseDto{
			Items:             roleResponse,
			PaginationOptions: &options,
		}

		return response, nil
	} else {
		return nil, nil
	}

}
