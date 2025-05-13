package usecase

import (
	"errors"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	userQuery "github.com/zercle/kku-qf-services/pkg/models/query-model"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u roleUsecase) DeleteRoleByID(roleID uint) (err error) {
	query := query.RoleQueryEntity{
		ID: &roleID,
	}

	err = u.CommonRepository.GetFirst(&query)

	if err != nil {
		return errors.New("Role not found")
	}

	return helper.ExecuteTransaction(u.CommonRepository, 
		u.DeleteRoleByIDTransaction(&query), 
		u.DeleteMapCoreRoles(&query),
		u.DeleteMapFacultiesRoles(&query),
		u.DeleteMapProgramsRoles(&query),
		u.ChangeUserRoleToDefault(&query))
}

func (u roleUsecase) DeleteRoleByIDTransaction(query *query.RoleQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, &query)
	}
}

func (u roleUsecase) DeleteMapCoreRoles(role *query.RoleQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		MapCoreRolesStatement := query.MapCoreRolesQueryEntity{
			RoleID: role.ID,
		}
		return u.CommonRepository.Delete(tx, &MapCoreRolesStatement)
	}
}

func (u roleUsecase) DeleteMapFacultiesRoles(role *query.RoleQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		MapCoreRolesStatement := query.MapFacultiesRolesQueryEntity{
			RoleID: role.ID,
		}
		return u.CommonRepository.Delete(tx, &MapCoreRolesStatement)
	}
}

func (u roleUsecase) DeleteMapProgramsRoles(role *query.RoleQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		MapCoreRolesStatement := query.MapProgramsRolesQueryEntity{
			RoleID: role.ID,
		}
		return u.CommonRepository.Delete(tx, &MapCoreRolesStatement)
	}
}

func (u roleUsecase) ChangeUserRoleToDefault(role *query.RoleQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		userStatement := userQuery.UserQueryEntity{
			CurrentRoleID: role.ID,
		}
		userUpdate := userQuery.UserQueryEntity{
			CurrentRoleID: pointer.ToUint(0),
		}
		return u.CommonRepository.Update(tx, userStatement, &userUpdate)
	}
}
