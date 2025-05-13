package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"

	"gorm.io/gorm"
)

func (u roleUsecase) DeleteSettingUserFacultyRole(userID *uuid.UUID, roleID, facultyID *uint) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteSettingUserFacultyRoleTransaction(userID, roleID, facultyID))
}

func (u roleUsecase) DeleteSettingUserFacultyRoleTransaction(userID *uuid.UUID, roleID, facultyID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		return u.CommonRepository.Delete(tx, &query.MapFacultiesRolesQueryEntity{
			UserID:    userID,
			RoleID:    roleID,
			FacultyID: facultyID,
		})
	}
}
