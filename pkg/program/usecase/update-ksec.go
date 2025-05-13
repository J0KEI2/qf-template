package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdateKsec(request dto.KsecDetail, ksecId uint) (result *dto.KsecDetail, err error) {
	ksecStatement := query.ProgramKsecDetailQueryEntity{
		ID: &ksecId,
	}
	ksecUpdate := query.ProgramKsecDetailQueryEntity{
		Order:     request.Order,
		Type:      request.Type,
		Detail:    request.Detail,
		IsChecked: request.IsChecked,
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdateKsecTransaction(&ksecStatement, &ksecUpdate))
	if err != nil {
		return nil, err
	}

	result = &dto.KsecDetail{
		ID:        ksecUpdate.ID,
		Order:     ksecUpdate.Order,
		Type:      ksecUpdate.Type,
		Detail:    ksecUpdate.Detail,
		IsChecked: ksecUpdate.IsChecked,
		CreatedAt: ksecUpdate.CreatedAt,
		UpdatedAt: ksecUpdate.UpdatedAt,
	}
	return
}

func (u programUsecase) UpdateKsecTransaction(statement *query.ProgramKsecDetailQueryEntity, update *query.ProgramKsecDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
