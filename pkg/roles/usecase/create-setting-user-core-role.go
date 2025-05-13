package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u roleUsecase) CreateSettingUserCoreRole(request dto.CreateSettingUserCoreRoleRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateSettingUserCoreRoleTransaction(request))
}

func (u roleUsecase) CreateSettingUserCoreRoleTransaction(request dto.CreateSettingUserCoreRoleRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		createBody := query.MapCoreRolesQueryEntity{
			UserID: request.UserID,
			RoleID: request.RoleID,
		}

		if err = u.CommonRepository.Create(tx, &createBody); err != nil {
			return err
		}

		return nil
	}
}
