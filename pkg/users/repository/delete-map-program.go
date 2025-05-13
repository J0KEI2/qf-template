package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *userRepository) DeleteMapProgramByUID(uid string) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn.Begin()
	defer dbTx.Rollback()
	dbTx = dbTx.Model(&migrateModels.MapUserProgram{})
	dbTx = dbTx.Where("user_uid = ?", uid)

	if err = dbTx.Delete(&migrateModels.MapUserProgram{}).Error; err != nil {
		return
	}

	if err = dbTx.Commit().Error; err != nil {
		return
	}

	return
}
