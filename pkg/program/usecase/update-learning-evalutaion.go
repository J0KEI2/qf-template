package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdateLearningEvaluation(request dto.LearningEvaluation, learningEvaluationId uint) (result *dto.LearningEvaluation, err error) {
	ksecStatement := query.ProgramPLOLearningEvaluationQueryEntity{
		ID: &learningEvaluationId,
	}
	ksecUpdate := query.ProgramPLOLearningEvaluationQueryEntity{
		Order:  request.Order,
		PloID:  request.PloID,
		Key:    request.Key,
		Detail: request.Detail,
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdateLearningEvaluationTransaction(&ksecStatement, &ksecUpdate))
	if err != nil {
		return nil, err
	}

	result = &dto.LearningEvaluation{
		ID:        ksecUpdate.ID,
		Order:     ksecUpdate.Order,
		PloID:     ksecUpdate.PloID,
		Key:       ksecUpdate.Key,
		Detail:    ksecUpdate.Detail,
		CreatedAt: ksecUpdate.CreatedAt,
		UpdatedAt: ksecUpdate.UpdatedAt,
	}
	return
}

func (u programUsecase) UpdateLearningEvaluationTransaction(statement *query.ProgramPLOLearningEvaluationQueryEntity, update *query.ProgramPLOLearningEvaluationQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
