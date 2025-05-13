package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdateLearningSolution(request dto.LearningSolution, learningSolutionId uint) (result *dto.LearningSolution, err error) {
	ksecStatement := query.ProgramPLOLearningSolutionQueryEntity{
		ID: &learningSolutionId,
	}
	ksecUpdate := query.ProgramPLOLearningSolutionQueryEntity{
		Order:  request.Order,
		PloID:  request.PloID,
		Key:    request.Key,
		Detail: request.Detail,
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdateLearningSolutionTransaction(&ksecStatement, &ksecUpdate))
	if err != nil {
		return nil, err
	}

	result = &dto.LearningSolution{
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

func (u programUsecase) UpdateLearningSolutionTransaction(statement *query.ProgramPLOLearningSolutionQueryEntity, update *query.ProgramPLOLearningSolutionQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
