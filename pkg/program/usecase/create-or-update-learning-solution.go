package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdatLearningSolution(learningSolutions []dto.CreateOrUpdateLearningSolutionRequestDto, ploID uint) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.createOrUpdateLearningSolutionTransaction(learningSolutions, ploID))
}

func (u programUsecase) createOrUpdateLearningSolutionTransaction(learningSolutions []dto.CreateOrUpdateLearningSolutionRequestDto, ploID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, learningSolution := range learningSolutions {
			queryPlo := query.ProgramPLOLearningSolutionQueryEntity{
				ID: learningSolution.ID,
			}
			update := query.ProgramPLOLearningSolutionQueryEntity{
				Order:  learningSolution.Order,
				PloID:  &ploID,
				Key:    learningSolution.Key,
				Detail: learningSolution.Detail,
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
