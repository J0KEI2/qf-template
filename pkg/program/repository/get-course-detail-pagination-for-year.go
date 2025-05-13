package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/common/repository"
	"github.com/zercle/kku-qf-services/pkg/models"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (r programRepository) GetCourseDetailPaginationForYear(paginationOptions *models.PaginationOptions, planID uint) (record []query.ProgramCourseDetailQueryEntity, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}
	dbTx := r.MainDbConn

	db := dbTx.Model(query.ProgramCourseDetailQueryEntity{})
	db.Where("year_and_semester_id IS NULL")
	db.Where("program_sub_plan_id = ?", planID)
	if paginationOptions.Search != nil {
		search := "%" + *paginationOptions.Search + "%"
		db.Where("(course_code LIKE ? OR course_name_th LIKE ? OR course_name_en LIKE ?)", search, search, search)
	}

	db.Count(paginationOptions.Total)
	db.Scopes(repository.Paginate(paginationOptions))

	record = make([]query.ProgramCourseDetailQueryEntity, 0)
	if err = db.Find(&record).Error; err != nil {
		return nil, err
	}
	return
}
