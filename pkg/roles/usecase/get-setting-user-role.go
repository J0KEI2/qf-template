package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	queryRoleAndPermission "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u roleUsecase) GetSettingUserFacultyRole(roleID, facultyID *uint, options models.PaginationOptions) (result *dto.GetSettingUserRoleResponseDto, err error) {
	if *roleID != 0 && *facultyID != 0 {
		roleResponse := []dto.SettingUserRoleDto{}
		userList := []queryRoleAndPermission.MapFacultiesRolesQueryEntity{}
		userQuery := queryRoleAndPermission.MapFacultiesRolesQueryEntity{
			FacultyID: facultyID,
			RoleID:    roleID,
		}

		options.SetSearchFields([]string{`"User".email`, `"User".name_th`, `"User".name_en`})

		if err = u.CommonRepository.GetList(&userQuery, &userList, &options, "Faculty", "Role", "User"); err != nil {
			return nil, err
		}

		for _, userItem := range userList {
			roleResponse = append(roleResponse, dto.SettingUserRoleDto{
				UserID:        userItem.User.UID,
				NameTH:        userItem.User.NameTH,
				NameEN:        userItem.User.NameEN,
				FacultyNameTH: userItem.Faculty.FacultyNameTH,
				FacultyNameEN: userItem.Faculty.FacultyNameEN,
				Email:         userItem.User.Email,
				CreatedAt:     userItem.CreatedAt,
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
