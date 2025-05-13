package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteKsec(id uint) (err error) {

	query := query.ProgramKsecDetailQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteKsecTransaction(&query))
}

func (u programUsecase) DeleteKsecTransaction(query *query.ProgramKsecDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
