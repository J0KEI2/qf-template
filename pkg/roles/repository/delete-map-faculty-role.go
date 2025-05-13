package repository

import (
	"fmt"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *roleRepository) DeleteSettingUserFacultyRole(userID uuid.UUID, roleID, facultyID uint) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn.Begin()
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&migrateModels.MapFacultiesRoles{})
	dbTx = dbTx.Where("user_id = ? AND role_id = ? AND faculty_id = ?", userID, roleID, facultyID)

	if err = dbTx.Delete(&migrateModels.MapFacultiesRoles{}).Error; err != nil {
		return
	}

	err = dbTx.Commit().Error

	return
}
