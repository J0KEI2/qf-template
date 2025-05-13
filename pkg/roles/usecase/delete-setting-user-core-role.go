package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"

	"gorm.io/gorm"
)

func (u roleUsecase) DeleteSettingUserCoreRole(userID *uuid.UUID, roleID *uint) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteSettingUserCoreRoleTransaction(userID, roleID))
}

func (u roleUsecase) DeleteSettingUserCoreRoleTransaction(userID *uuid.UUID, roleID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		if err = u.RoleRepository.DeleteSettingUserCoreRole(*userID, *roleID); err != nil {
			return err
		}

		return nil
	}
}
