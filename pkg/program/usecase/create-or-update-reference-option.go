package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	commonQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateReferenceOption(reference dto.ProgramReferenceDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateReferenceOptionTransaction(reference))
}

func (u programUsecase) CreateOrUpdateReferenceOptionTransaction(reference dto.ProgramReferenceDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryRef := commonQuery.ReferenceOptionQueryEntity{
			ID:   reference.ReferenceTypeID,
			Name: reference.ReferenceTypeName,
		}

		err = u.CommonRepository.GetFirstOrCreate(&queryRef)
		if err != nil {
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}

		return
	}
}
