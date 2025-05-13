package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	migrate_models "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *userRepository) GetUsers(criteria *models.UserFetchQuery) (users []models.UserFetchModel, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn.Model(&migrate_models.Users{})

	dbTx = dbTx.Order("created_at DESC")

	if criteria != nil {
		if criteria.UID != nil {
			dbTx = dbTx.Where("uid = ?", *criteria.UID)
		}
		if criteria.Email != nil {
			dbTx = dbTx.Where("email = ?", *criteria.Email)
		}
		if criteria.SSN != nil {
			dbTx = dbTx.Where("ssn = ?", *criteria.SSN)
		}
		if criteria.SystemPermissionUID != nil {
			dbTx = dbTx.Where("system_permission_uid = ?", *criteria.SystemPermissionUID)
		}
		if criteria.FacultyID != nil {
			dbTx = dbTx.Where("faculty_id = ?", *criteria.FacultyID)
		}
		if criteria.Type != nil {
			dbTx = dbTx.Where("type = ?", *criteria.Type)
		}
		if criteria.Status != nil {
			dbTx = dbTx.Where("status = ?", *criteria.Status)
		}

		if criteria.OrderBy != nil && criteria.Direction != nil {
			dbTx = dbTx.Order(fmt.Sprintf("%s %s", *criteria.OrderBy, *criteria.Direction))
		}

		if criteria.Limit != nil && criteria.Offset != nil {
			dbTx = dbTx.Limit(*criteria.Limit)
			dbTx = dbTx.Offset(*criteria.Offset)
		}
	}

	err = dbTx.Find(&users).Error

	return
}
