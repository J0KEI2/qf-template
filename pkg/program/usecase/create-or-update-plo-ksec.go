package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdatePlo(plo dto.CreateOrUpdatePLORequestDto, subPlanID uint) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.createOrUpdatePloTransaction(plo, subPlanID), u.createOrUpdateKsecTransaction(plo, subPlanID))
}

func (u programUsecase) createOrUpdatePloTransaction(plo dto.CreateOrUpdatePLORequestDto, subPlanID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryPlo := query.ProgramPloFormatQueryEntity{
			ID:           &plo.ID,
			ProgramSubPlanID: &subPlanID,
		}
		update := query.ProgramPloFormatQueryEntity{
			PLOFormat: plo.PLOFormat,
		}
		err = u.CommonRepository.Update(tx, queryPlo, &update)
		if err != nil {
			return err
		}
		return u.ProgramRepository.CreateOrUpdatePLODetail(tx, plo.PLODetails, &plo.ID, nil)
	}
}

func (u programUsecase) createOrUpdateKsecTransaction(plo dto.CreateOrUpdatePLORequestDto, subPlanID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		return u.ProgramRepository.CreateOrUpdateKsecDetail(tx, plo.Ksec, &subPlanID)
	}
}
