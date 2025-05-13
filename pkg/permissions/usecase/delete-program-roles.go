package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) DeleteProgramRoles(mapId uint) (err error) {

	query := query.MapProgramsRolesQueryEntity{
		ID: &mapId,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, deleteMapProgramsRoleStatement(u, &query)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return err
	}
	return nil
}

func deleteMapProgramsRoleStatement(u *PermissionUsecase, queryDb *query.MapProgramsRolesQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, queryDb)
	}
}
