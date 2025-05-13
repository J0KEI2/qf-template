package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteMapPloKsec(id uint) (err error) {

	query := query.ProgramMapPloWithKsecQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteMapPloKsecTransaction(&query))
}

func (u programUsecase) DeleteMapPloKsecTransaction(query *query.ProgramMapPloWithKsecQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
