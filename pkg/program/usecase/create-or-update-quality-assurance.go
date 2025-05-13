package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateQualityAssurance(qa dto.CreateOrUpdateQualityAssuranceDto, programMainUID uuid.UUID) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateQualityAssuranceTransaction(qa, programMainUID))
}

func (u programUsecase) CreateOrUpdateQualityAssuranceTransaction(qa dto.CreateOrUpdateQualityAssuranceDto, programMainUID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		programMain := query.ProgramMainQueryEntity{
			ID: &(programMainUID),
		}
		err = u.CommonRepository.GetFirst(&programMain)

		if err != nil {
			return err
		}

		queryPlo := query.ProgramQualityAssuranceQueryEntity{
			ID: programMain.ProgramQualityAssuranceID,
		}

		update := query.ProgramQualityAssuranceQueryEntity{
			IsHescCheck:      qa.HESC.IsCheck,
			HescDescription:  qa.HESC.Description,
			IsAunQaCheck:     qa.AunQa.IsCheck,
			AunQaDescription: qa.AunQa.Description,
			IsAbetCheck:      qa.ABET.IsCheck,
			AbetDescription:  qa.ABET.Description,
			IsWfmeCheck:      qa.WFME.IsCheck,
			WfmeDescription:  qa.WFME.Description,
			IsAacsbCheck:     qa.AACSB.IsCheck,
			AacsbDescription: qa.AACSB.Description,
		}
		return u.CommonRepository.Update(tx, queryPlo, &update)
	}
}
