package repository

import (
	"fmt"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *roleRepository) DeleteSettingUserCoreRole(userID uuid.UUID, roleID uint) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn.Begin()
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&migrateModels.MapCoreRoles{})
	dbTx = dbTx.Where("user_id = ? AND role_id = ?", userID, roleID)

	if err = dbTx.Delete(&migrateModels.MapCoreRoles{}).Error; err != nil {
		return
	}

	err = dbTx.Commit().Error

	return
}
