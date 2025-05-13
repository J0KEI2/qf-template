package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	// migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

// TODO: create the new func don't forget to embed in domain repository interface
func (repo *courseRepository) GetMapCourseLecturer(id int) (response []dto.MapCourseLecturer, err error) {
	if repo.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := repo.MainDbConn
	dbTx = dbTx.Where("l.id = ?", id)
	dbTx = dbTx.Table("course_lecturers as l").
		Select("m.id, m.course_lecturer_id, m.course_lecturer_id")
	dbTx = dbTx.Joins("JOIN map_course_lecturers as m ON l.id = m.course_lecturer_id")

	err = dbTx.Find(&response).Error

	return
}
