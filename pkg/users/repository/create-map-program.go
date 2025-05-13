package repository

import (
	"fmt"

	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"

	helpers "github.com/zercle/gofiber-helpers"
)

func (r *userRepository) CreateMapProgram(newData migrateModels.MapUserProgram) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn.Begin()
	// if err rollback transaction
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&migrateModels.MapUserProgram{})
	if err = dbTx.Create(newData).Error; err != nil {
		return err
	}

	if err = dbTx.Commit().Error; err != nil {
		return err
	}

	return
}
