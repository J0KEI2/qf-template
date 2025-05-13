package repository

import (
	"fmt"
	"strings"

	"github.com/AlekSi/pointer"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (r *userRepository) GetUser(criteria *models.UserFetchQuery) (user *entity.UserFetchEntity, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	res := entity.UserQueryEntityWithRole{
		SSN: criteria.SSN,
	}

	dbTx := r.MainDbConn.Model(res)

	dbTx.Where(res).Joins("LEFT JOIN roles as r ON u.current_role_id = r.id")

	if err = dbTx.First(&res).Error; err != nil {
		return nil, err
	}

	nameTHList := append(strings.SplitN(pointer.GetString(res.NameTH), " ", 3), []string{"", ""}...)
	nameENList := append(strings.SplitN(pointer.GetString(res.NameEN), " ", 3), []string{"", ""}...)

	user = &entity.UserFetchEntity{
		CreatedAt:           *res.CreatedAt,
		LastAccessAt:        *res.LastAccessAt,
		UpdatedAt:           *res.UpdatedAt,
		Email:               *res.Email,
		Status:              *res.Status,
		SSN:                 *res.SSN,
		RoleNameTH:          *res.RoleNameTH,
		RoleNameEN:          *res.RoleNameEN,
		Type:                *res.Type,
		TitleTh:             nameTHList[0],
		FirstnameTh:         nameTHList[1],
		LastnameTH:          nameTHList[2],
		TitleEn:             nameENList[0],
		FirstnameEn:         nameENList[1],
		LastnameEn:          nameENList[2],
		FacultyID:           *res.FacultyID,
		CurrentRoleID:       *res.CurrentRoleID,
		SystemPermissionUID: *res.SystemPermissionUID,
		UID:                 *res.UID,
	}

	return
}
