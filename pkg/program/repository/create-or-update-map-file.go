package repository

import (
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateMapFile(tx *gorm.DB, fileId, id *uint) (err error) {
	update := query.MapFilesSystemQueryEntity{
		FileID:      fileId,
		ReferenceID: id,
	}
	if id != nil {
		if err = tx.Where(&query.MapFilesSystemQueryEntity{
			ReferenceID: id,
		}).Updates(&update).Error; err != nil {
			return err
		}
	} else {
		err = tx.Create(&update).Error
		if err != nil {
			return err
		}
	}

	return nil
}
