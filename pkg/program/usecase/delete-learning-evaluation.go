package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteLearningEvaluation(id uint) (err error) {

	query := query.ProgramPLOLearningEvaluationQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteLearningEvaluationTransaction(&query))
}

func (u programUsecase) DeleteLearningEvaluationTransaction(query *query.ProgramPLOLearningEvaluationQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
