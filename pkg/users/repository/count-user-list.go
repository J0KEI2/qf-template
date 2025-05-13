package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"gorm.io/gorm"
)

func (r *userRepository) CountUserList(criteria *models.UserFetchWithRelationQueryModel) (count *int64, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}
	totalRecord := new(int64)

	dbTx := r.MainDbConn
	dbTx = getCountUserListBDQuery(dbTx, criteria)

	dbTx = dbTx.Table("users u").
		Select("u.uid, u.email, u.ssn, u.name_th, u.name_en, u.system_permission_uid, u.faculty_id, u.type, u.status, u.created_at, u.updated_at, u.last_access_at, r.role_name_th, r.role_name_en,f.faculty_name_th,f.faculty_name_en").
		Joins("LEFT JOIN roles as r ON u.current_role_id = r.id").
		Joins("LEFT JOIN common_faculties as f ON u.faculty_id = f.id")

	err = dbTx.Count(totalRecord).Error
	if err != nil {
		return nil, err
	}

	return totalRecord, nil
}

func getCountUserListBDQuery(dbTx *gorm.DB, criteria *models.UserFetchWithRelationQueryModel) *gorm.DB {
	dbTx = dbTx.Order("created_at DESC")
	if criteria != nil {

		if criteria.Name != nil {
			dbTx = dbTx.Where("u.name_th LIKE ? OR u.name_en LIKE ?", "%"+*criteria.Name+"%", "%"+*criteria.Name+"%")
		}

		if criteria.FacultyNameTH != nil {
			dbTx = dbTx.Where("f.faculty_name_th LIKE ?", "%"+*criteria.FacultyNameTH+"%")
		}

		if criteria.RoleNameEN != nil {
			dbTx = dbTx.Where("r.role_name_en = ?", *criteria.RoleNameEN)
		}

		if criteria.UID != nil {
			dbTx = dbTx.Where("u.uid = ?", *criteria.UID)
		}
		if criteria.Email != nil {
			dbTx = dbTx.Where("u.email = ?", *criteria.Email)
		}
		if criteria.SSN != nil {
			dbTx = dbTx.Where("u.ssn = ?", *criteria.SSN)
		}
		if criteria.SystemPermissionUID != nil {
			dbTx = dbTx.Where("u.system_permission_uid = ?", *criteria.SystemPermissionUID)
		}
		if criteria.FacultyID != nil {
			dbTx = dbTx.Where("u.faculty_id = ?", *criteria.FacultyID)
		}
		if criteria.Type != nil {
			dbTx = dbTx.Where("u.type = ?", *criteria.Type)
		}
		if criteria.Status != nil {
			dbTx = dbTx.Where("u.status = ?", *criteria.Status)
		}

	}

	return dbTx
}
