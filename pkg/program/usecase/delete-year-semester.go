package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteYearAndSemester(id uint) (err error) {

	query := query.ProgramYearAndSemesterQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteYearAndSemesterTransaction(&query))
}

func (u programUsecase) DeleteYearAndSemesterTransaction(query *query.ProgramYearAndSemesterQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
