package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) CreateCoreRoles(request dto.CreateMapCoreRolesRequestDto) (result *dto.MapCoreRoles, err error) {

	queryMap := query.MapCoreRolesQueryEntity{
		UserID: &request.UserID,
		RoleID: &request.RoleID,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, createMapCoreRoleStatement(u, &queryMap)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	queryRole := query.RoleQueryEntity{
		ID: &request.RoleID,
	}

	u.CommonRepository.GetFirst(&queryRole)

	result = &dto.MapCoreRoles{
		ID:         queryMap.ID,
		UserID:     queryMap.UserID,
		RoleID:     queryMap.RoleID,
		RoleNameTH: queryRole.RoleNameTH,
		RoleNameEN: queryRole.RoleNameEN,
	}
	return result, nil
}

func createMapCoreRoleStatement(u *PermissionUsecase, queryDb *query.MapCoreRolesQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Create(tx, queryDb)
	}
}
