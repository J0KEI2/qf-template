package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdateMapPloKsec(request dto.ProgramMapPloWithKsec, ksecId uint) (result *dto.ProgramMapPloWithKsec, err error) {
	mapPloKsecStatement := query.ProgramMapPloWithKsecQueryEntity{
		ID: &ksecId,
	}
	mapPloKsecUpdate := query.ProgramMapPloWithKsecQueryEntity{
		PloID:  request.PloID,
		KsecID: request.KsecID,
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdateMapPloWithKsecTransaction(&mapPloKsecStatement, &mapPloKsecUpdate))
	if err != nil {
		return nil, err
	}

	result = &dto.ProgramMapPloWithKsec{
		ID:        mapPloKsecUpdate.ID,
		PloID:     mapPloKsecUpdate.PloID,
		KsecID:    mapPloKsecUpdate.KsecID,
		CreatedAt: mapPloKsecUpdate.CreatedAt,
		UpdatedAt: mapPloKsecUpdate.UpdatedAt,
	}
	return
}

func (u programUsecase) UpdateMapPloWithKsecTransaction(statement *query.ProgramMapPloWithKsecQueryEntity, update *query.ProgramMapPloWithKsecQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
