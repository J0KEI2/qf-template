package repository

import (
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) UpdateRefNilOption(tx *gorm.DB, id *uint) (err error) {
	if err = tx.Model(&programQuery.ProgramReferenceQueryEntity{}).Where("id = ?", *id).UpdateColumn("reference_type_id", gorm.Expr("NULL")).Error; err != nil {
		return err
	}

	return nil
}
