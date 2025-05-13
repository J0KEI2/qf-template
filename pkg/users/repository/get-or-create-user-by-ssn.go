package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	migrate_models "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"github.com/zercle/kku-qf-services/pkg/utils"
)

func (r *userRepository) GetOrCreateBySSN(userCreateQuery models.UserCreateQuery) (*entity.UserFetchEntity, error) {
	if r.MainDbConn == nil {
		err := fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	dbTx := r.MainDbConn.Begin()
	// if err rollback transaction
	defer dbTx.Rollback()

	// create user
	userDbTx := dbTx.Model(&migrate_models.Users{})

	user := &migrate_models.Users{
		SSN: userCreateQuery.SSN,
	}

	nameTH := strings.Join([]string{userCreateQuery.TitleTh, userCreateQuery.FirstnameTh, userCreateQuery.LastnameTh}, " ")
	nameEN := strings.Join([]string{userCreateQuery.TitleEn, userCreateQuery.FirstnameEn, userCreateQuery.LastnameEn}, " ")

	if err := userDbTx.Where(migrate_models.Users{SSN: userCreateQuery.SSN}).Attrs(
		migrate_models.Users{
			UID:                 uuid.New(),
			Email:               userCreateQuery.Email,
			SSN:                 userCreateQuery.SSN,
			SystemPermissionUID: userCreateQuery.SystemPermissionUID,
			NameTH:              nameTH,
			NameEN:              nameEN,
			FacultyID:           userCreateQuery.FacultyID,
			Type:                userCreateQuery.Type,
			Status:              "ACTIVE",
		}).FirstOrCreate(&user).Error; err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	if err := dbTx.Commit().Error; err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	userResult := &entity.UserFetchEntity{
		UID:                 user.UID,
		Email:               user.Email,
		SSN:                 utils.GetStringFromPointer(user.SSN),
		SystemPermissionUID: user.SystemPermissionUID,
		FacultyID:           user.FacultyID,
		Type:                string(user.Type),
		Status:              string(user.Status),
		TitleTh:             userCreateQuery.TitleTh,
		FirstnameTh:         userCreateQuery.FirstnameTh,
		LastnameTH:          userCreateQuery.LastnameTh,
		TitleEn:             userCreateQuery.TitleTh,
		FirstnameEn:         userCreateQuery.FirstnameEn,
		LastnameEn:          userCreateQuery.LastnameEn,
		CurrentRoleID:       user.CurrentRoleID,
	}

	return userResult, nil
}
