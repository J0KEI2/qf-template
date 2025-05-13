package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteMapFileSystem(id uint) (err error) {

	query := query.MapFilesSystemQueryEntity{
		FileID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteMapFileSystemTransaction(&query))
}

func (u programUsecase) DeleteMapFileSystemTransaction(query *query.MapFilesSystemQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
