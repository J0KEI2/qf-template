package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateLearningEvaluation(learningEvaluation []dto.CreateOrUpdateLearningEvaluationRequestDto, ploID uint) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.createOrUpdateLearningEvaluationTransaction(learningEvaluation, ploID))
}

func (u programUsecase) createOrUpdateLearningEvaluationTransaction(LearningEvaluations []dto.CreateOrUpdateLearningEvaluationRequestDto, ploID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, LearningEvaluation := range LearningEvaluations {
			queryPlo := query.ProgramPLOLearningEvaluationQueryEntity{
				ID: LearningEvaluation.ID,
			}
			update := query.ProgramPLOLearningEvaluationQueryEntity{
				Order:  LearningEvaluation.Order,
				PloID:  &ploID,
				Key:    LearningEvaluation.Key,
				Detail: LearningEvaluation.Detail,
			}
			err = u.CommonRepository.Update(tx, queryPlo, &update)
			if err != nil {
				if err != gorm.ErrRecordNotFound {
					err = u.CommonRepository.Create(tx, &update)
				}
				if err != nil {
					return err
				}
			}
		}
		return
	}
}
