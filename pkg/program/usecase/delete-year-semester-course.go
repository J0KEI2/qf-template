package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteYearCourse(id uint) (err error) {
	statement := query.ProgramCourseDetailQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteYearCourseTransaction(&statement))
}

func (u programUsecase) DeleteYearCourseTransaction(statement *query.ProgramCourseDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		deleteYearAndSemester := query.ProgramCourseDetailQueryEntity{
			YearAndSemesterID: nil,
		}
		return u.CommonRepository.Update(tx, statement, &deleteYearAndSemester)
	}
}
