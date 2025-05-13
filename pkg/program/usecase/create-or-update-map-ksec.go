package usecase

import (
	"log"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateMapPloWithKsec(KsecRequest dto.KsecRequestDto, ploID uint) (err error) {
	KsecDetailArray := KsecRequest.ToKsecDetail()
	log.Println(KsecDetailArray)
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateMapPloWithKsecTransaction(KsecDetailArray, ploID))
}

func (u programUsecase) CreateOrUpdateMapPloWithKsecTransaction(KsecDetails []dto.KsecDetail, ploID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, KsecDetail := range KsecDetails {
			queryUpdate := query.ProgramMapPloWithKsecQueryEntity{
				ID: KsecDetail.MapPLOID,
			}
			update := query.ProgramMapPloWithKsecQueryEntity{
				PloID:  &ploID,
				KsecID: KsecDetail.ID,
			}
			log.Println(KsecDetail.ID, KsecDetail.IsChecked)
			if pointer.GetBool(KsecDetail.IsChecked) {
				if KsecDetail.MapPLOID == nil {
					err = u.CommonRepository.Create(tx, &update)
				}
			} else {
				if KsecDetail.MapPLOID != nil {
					err = u.CommonRepository.Delete(tx, &queryUpdate)
				}
			}
			if err != nil {
				return err
			}
		}

		return
	}
}
