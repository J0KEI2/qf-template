package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	MGModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *userDetailRepository) GetAllHrKeyLecturers() (users []models.CronUpdateLecturer, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}
	// selectColumn := "users.hr_key, userDetail.uid, userDetail.title_th, userDetail.firstname_th, userDetail.lastname_th, userDetail.title_en, userDetail.title_th, user.status, userDetail.position, userDetail.education_back_ground"
	selectColumn := "users.hr_key, userDetails.uid"
	dbTx := r.MainDbConn.Model(&MGModels.Users{})
	err = dbTx.Select(selectColumn).Joins("left join userDetails on userDetails.user_uid = users.uid").Scan(&users).Error

	return
}
