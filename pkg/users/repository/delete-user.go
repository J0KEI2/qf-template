package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *userRepository) DeleteUser(userID string) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn.Begin()
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&migrateModels.Users{})
	dbTx = dbTx.Where("uid = ?", userID)

	if err = dbTx.Delete(&migrateModels.Users{}).Error; err != nil {
		return
	}

	err = dbTx.Commit().Error

	return
}
