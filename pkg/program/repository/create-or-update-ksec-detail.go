package repository

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateKsecDetail(tx *gorm.DB, ksec dto.KsecRequestDto, subPlanID *uint) (err error) {
	ksecDetails := ksec.ToKsecDetail()
	for _, ksecDetail := range ksecDetails {
		updateQuery := query.ProgramKsecDetailQueryEntity{
			ID: ksecDetail.ID,
		}
		update := query.ProgramKsecDetailQueryEntity{
			ProgramSubPlanID: subPlanID,
			Order:        ksecDetail.Order,
			Type:         ksecDetail.Type,
			Detail:       ksecDetail.Detail,
		}
		if err = tx.Where(updateQuery).Updates(&update).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				err = tx.Create(&update).Error
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r programRepository) CreateKsecDetail(tx *gorm.DB, ksec dto.KsecRequestDto, subPlanID *uint) (err error) {
	ksecDetails := ksec.ToKsecDetail()
	for _, ksecDetail := range ksecDetails {
		update := query.ProgramKsecDetailQueryEntity{
			ProgramSubPlanID: subPlanID,
			Order:        ksecDetail.Order,
			Type:         ksecDetail.Type,
			Detail:       ksecDetail.Detail,
		}
		err = tx.Create(&update).Error
		if err != nil {
			return err
		}
	}
	return nil
}
