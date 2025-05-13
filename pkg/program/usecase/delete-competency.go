package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteCompetency(competencyID uint) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.deleteCompetencyAction(competencyID))
}

func (u *programUsecase) deleteCompetencyAction(competencyID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		competencyQuery := query.ProgramCompetencyQueryEntity{
			ID: &competencyID,
		}

		err = u.CommonRepository.Delete(tx, &competencyQuery)
		if err != nil {
			return err
		}

		return
	}
}
