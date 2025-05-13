package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	// migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

// TODO: create the new func don't forget to embed in domain repository interface
func (repo *qf4Repository) GetMapQF4Lecturer(id int) (response []dto.MapQF4Lecturer, err error) {
	if repo.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := repo.MainDbConn
	dbTx = dbTx.Where("l.id = ?", id)
	dbTx = dbTx.Table("qf4_lecturers as l").
		Select("m.id, m.qf4_lecturer_id, m.course_lecturer_id")
	dbTx = dbTx.Joins("JOIN map_qf4_lecturers as m ON l.id = m.qf4_lecturer_id")

	err = dbTx.Find(&response).Error

	return
}
