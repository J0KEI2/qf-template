package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteLearningSolution(id uint) (err error) {

	query := query.ProgramPLOLearningSolutionQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteLearningSolutionTransaction(&query))
}

func (u programUsecase) DeleteLearningSolutionTransaction(query *query.ProgramPLOLearningSolutionQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
