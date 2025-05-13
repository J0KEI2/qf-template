package usecase

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u roleUsecase) UpdateUserCurrentRole(userID uuid.UUID, roleID uint) (err error) {
	userStatement := query.UserQueryEntity{
		UID: &userID,
	}

	userUpdate := query.UserQueryEntity{
		CurrentRoleID: &roleID,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateUserCurrentRoleTransaction(userStatement, &userUpdate))
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return err
	}

	return nil
}

func (u roleUsecase) updateUserCurrentRoleTransaction(statement query.UserQueryEntity, update *query.UserQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
