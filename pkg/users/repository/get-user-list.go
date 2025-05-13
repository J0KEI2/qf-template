package repository

import (
	"fmt"
	"strings"

	"github.com/AlekSi/pointer"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"gorm.io/gorm"
)

func (r *userRepository) GetUserList(criteria *models.UserFetchWithRelationQueryModel, options *models.PaginationOptions) (users []models.UserFetchWithRelationModel, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn
	dbTx = getUserListBDQuery(dbTx, criteria)
	dbTx = dbTx.Table("users as u").
		Select("u.uid, u.email, u.name_th, u.name_en, u.ssn, u.system_permission_uid, u.faculty_id, u.type, u.status, u.created_at, u.updated_at, u.last_access_at, r.role_name_th, r.role_name_en,f.faculty_name_th,f.faculty_name_en").
		Joins("LEFT JOIN roles as r ON u.current_role_id = r.id").
		Joins("LEFT JOIN common_faculties as f ON u.faculty_id = f.id").
		Where("u.deleted_at is null")
	if options != nil && options.Search != nil {
		searchParam := "%" + strings.ToLower(pointer.GetString(options.Search)) + "%"
		dbTx.Where("LOWER(u.email) LIKE ? OR LOWER(u.name_th) LIKE ? OR LOWER(u.name_en) LIKE ?", searchParam, searchParam, searchParam)
	}
	err = dbTx.Find(&users).Error

	return
}

func getUserListBDQuery(dbTx *gorm.DB, criteria *models.UserFetchWithRelationQueryModel) *gorm.DB {
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
		if email := pointer.GetString(criteria.Email); email != "" {
			email = "%" + email + "%"
			dbTx = dbTx.Where("u.email LIKE ?", email)
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

		if criteria.OrderBy != nil && criteria.Direction != nil {
			dbTx = dbTx.Order(fmt.Sprintf("%s %s", mappingOderBy(*criteria.OrderBy), *criteria.Direction))
		}

		if criteria.Limit != nil && criteria.Offset != nil {
			dbTx = dbTx.Limit(*criteria.Limit)
			dbTx = dbTx.Offset(*criteria.Offset)
		}
	}

	return dbTx
}

func mappingOderBy(oderBy string) string {
	switch oderBy {
	case entity.EmailName:
		return fmt.Sprintf(`u.%s`, entity.EmailName)
	case entity.FirstNameTH:
		return fmt.Sprintf(`l.%s`, entity.FirstNameTH)
	case entity.FirstNameEN:
		return fmt.Sprintf(`l.%s`, entity.FirstNameEN)
	case entity.FacultyNameTH:
		return fmt.Sprintf(`f.%s`, entity.FacultyNameTH)
	case entity.RoleNameEN:
		return fmt.Sprintf(`r.%s`, entity.RoleNameEN)
	case "created_at":
		return "u.created_at"
	case "updated_at":
		return "u.updated_at"
	}
	return "u.created_at"
}
