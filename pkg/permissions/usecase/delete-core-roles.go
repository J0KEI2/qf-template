package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) DeleteCoreRoles(mapId uint) (err error) {

	query := query.MapCoreRolesQueryEntity{
		ID: &mapId,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, deleteMapCoreRoleStatement(u, &query)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return err
	}
	return nil
}

func deleteMapCoreRoleStatement(u *PermissionUsecase, queryDb *query.MapCoreRolesQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, queryDb)
	}
}
