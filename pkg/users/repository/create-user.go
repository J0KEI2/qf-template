package repository

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	migrate_models "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"github.com/zercle/kku-qf-services/pkg/utils"
)

func (r *userRepository) CreateUser(createUserDto models.UserCreateQuery) (userResult *entity.UserFetchEntity, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn.Begin()
	// if err rollback transaction
	defer dbTx.Rollback()

	nameTH := strings.Join([]string{createUserDto.TitleTh, createUserDto.FirstnameTh, createUserDto.LastnameTh}, " ")
	nameEN := strings.Join([]string{createUserDto.TitleEn, createUserDto.FirstnameEn, createUserDto.LastnameEn}, " ")

	//create user database
	userDbTx := dbTx.Model(&migrate_models.Users{})

	user := &migrate_models.Users{
		UID:                 uuid.New(),
		Email:               createUserDto.Email,
		SSN:                 createUserDto.SSN,
		SystemPermissionUID: createUserDto.SystemPermissionUID,
		NameTH:              nameTH,
		NameEN:              nameEN,
		FacultyID:           createUserDto.FacultyID,
		Type:                createUserDto.Type,
		Status:              "ACTIVE",
	}

	if err = userDbTx.Create(user).Error; err != nil {
		return nil, err
	}

	if err = dbTx.Commit().Error; err != nil {
		return nil, err
	}

	userResult = &entity.UserFetchEntity{
		UID:                 user.UID,
		Email:               user.Email,
		SSN:                 utils.GetStringFromPointer(user.SSN),
		SystemPermissionUID: user.SystemPermissionUID,
		FacultyID:           user.FacultyID,
		Type:                string(user.Type),
		Status:              string(user.Status),
		TitleTh:             createUserDto.TitleTh,
		FirstnameTh:         createUserDto.FirstnameTh,
		LastnameTH:          createUserDto.LastnameTh,
		TitleEn:             createUserDto.TitleTh,
		FirstnameEn:         createUserDto.FirstnameEn,
		LastnameEn:          createUserDto.LastnameEn,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
	}

	return userResult, err
}
