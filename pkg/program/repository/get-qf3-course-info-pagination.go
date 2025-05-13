package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/common/repository"
	"github.com/zercle/kku-qf-services/pkg/models"
	courseQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (r programRepository) GetCourseInfoPagnation(paginationOptions *models.PaginationOptions) (record []courseQuery.CourseMainQueryEntity, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}
	dbTx := r.MainDbConn

	db := dbTx.Model(courseQuery.CourseMainQueryEntity{})
	if paginationOptions.Search != nil {
		search := "%" + *paginationOptions.Search + "%"
		db.Joins(`JOIN course_infos ON courses.course_info_id = course_infos.id`)
		db.Where(`(course_infos.course_code LIKE ? 
			OR course_infos.course_name_th LIKE ? 
			OR course_infos.course_name_en LIKE ?)`,
			search, search, search)
	}
	db.Preload("CourseInfo")

	db.Count(paginationOptions.Total)
	db.Scopes(repository.Paginate(paginationOptions))

	record = make([]courseQuery.CourseMainQueryEntity, 0)
	if err = db.Find(&record).Error; err != nil {
		return nil, err
	}
	return
}
