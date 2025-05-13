package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateCompetency(competency dto.ProgramCompetencyRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateCompetencyTransaction(competency))
}

func (u programUsecase) CreateOrUpdateCompetencyTransaction(competency dto.ProgramCompetencyRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, competencyData := range competency.ProgramCompetencyList {
			queryCompetency := query.ProgramCompetencyQueryEntity{
				ID: competencyData.ID,
			}
			update := query.ProgramCompetencyQueryEntity{
				ID:                 competencyData.ID,
				ProgramMainID:      &competency.ProgramMainID,
				Order:              competencyData.Order,
				SpecificCompetency: competencyData.SpecificCompetency,
				GenericCompetency:  competencyData.GenericCompetency,
			}
			if err = u.CommonRepository.Update(tx, queryCompetency, &update); err != nil {
				if err != gorm.ErrRecordNotFound {
					err = u.CommonRepository.Create(tx, &update)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}
		}

		return
	}
}
