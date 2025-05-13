package repository

import (
	"encoding/json"
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *userRepository) UpdateRoleUser(uid string, patchRequest models.UpdateUserRoleRequest) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	// Mock
	// position = enums.PROFESSOR
	educationBG := entity.EducationBackGround{
		Qualification:    "A",
		Department:       "โรงพยาบาลศรีนครินทร์",
		InstituteName:    "B",
		InstituteCountry: "C",
		GraduateYear:     2020,
	}

	educationBgByteData, err := json.Marshal(educationBG)
	if err != nil {
		return err
	}

	positionByteData, err := json.Marshal(patchRequest.Position)
	if err != nil {
		return err
	}

	educationBgStr := string(educationBgByteData)
	positionStr := string(positionByteData)
	newUserDetailData := migrateModels.UserDetail{
		UserUID:             patchRequest.UserUID,
		EducationBackGround: &educationBgStr,
		Position:            &positionStr,
		// YearOfAcceptingPosition: patchRequest.YearOfAcceptingPosition, // wait to add this column
	}

	dbTx := r.MainDbConn.Begin()
	defer dbTx.Rollback()

	dbTx = dbTx.Model(&migrateModels.UserDetail{})
	dbTx = dbTx.Where("user_uid = ?", uid)

	if err = dbTx.Updates(newUserDetailData).Error; err != nil {
		return
	}

	if err = dbTx.Commit().Error; err != nil {
		return
	}

	return
}
