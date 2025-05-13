package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u roleUsecase) CreateSettingUserFacultyRole(request dto.CreateSettingUserFacultyRoleRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateSettingUserFacultyRoleTransaction(request))
}

func (u roleUsecase) CreateSettingUserFacultyRoleTransaction(request dto.CreateSettingUserFacultyRoleRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		createBody := query.MapFacultiesRolesQueryEntity{
			UserID:    request.UserID,
			RoleID:    request.RoleID,
			FacultyID: request.FacultyID,
		}

		if err = u.CommonRepository.Create(tx, &createBody); err != nil {
			return err
		}

		return nil
	}
}
