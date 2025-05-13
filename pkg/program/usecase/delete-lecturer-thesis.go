package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteLecturerThesis(id uint) (err error) {

	query := query.ProgramThesisLecturerQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteLecturerThesisTransaction(&query))
}

func (u programUsecase) DeleteLecturerThesisTransaction(query *query.ProgramThesisLecturerQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}
