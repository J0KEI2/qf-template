package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteStructure(id uint) (err error) {

	query := query.ProgramStructureDetailQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteStructureTransaction(&query))
}

func (u programUsecase) DeleteStructureTransaction(query *query.ProgramStructureDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
