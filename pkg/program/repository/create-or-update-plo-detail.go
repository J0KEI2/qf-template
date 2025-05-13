package repository

import (
	"log"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdatePLODetail(tx *gorm.DB, ploDetails []dto.ProgramPLODetailDto, ploId *uint, parentID *uint) (err error) {
	for _, ploDetail := range ploDetails {
		statement := query.ProgramPloQueryEntity{
			ID:                 ploDetail.ID,
		}
		update := query.ProgramPloQueryEntity{
			ProgramPloFormatID: ploId,
			Order:              ploDetail.Order,
			ParentID:           parentID,
			PLOPrefix:          ploDetail.PLOPrefix,
			PLODetail:          ploDetail.PLODetail,
		}
		if err = tx.Where(statement).Updates(&update).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				err = tx.Create(&update).Error
			}
			if err != nil {
				return err
			}
		}
		if err = r.CreateOrUpdatePLODetail(tx, ploDetail.Children, ploId, update.ID); err != nil {
			return err
		}
	}
	return nil
}

func (r programRepository) CreatePLODetail(tx *gorm.DB, ploDetails []dto.ProgramPLODetailDto, ploId *uint, parentID *uint, ksecDetails []dto.MapKsecList) error {
	for index, ploDetail := range ploDetails {
		update := query.ProgramPloQueryEntity{
			ID:                 ploDetail.ID,
			ProgramPloFormatID: ploId,
			Order:              ploDetail.Order,
			ParentID:           parentID,
			PLOPrefix:          ploDetail.PLOPrefix,
			PLODetail:          ploDetail.PLODetail,
		}
		if err := tx.Create(&update).Error; err != nil {
			return err
		}

		ksecDetailArray := ksecDetails[index].KsecList.ToKsecDetail()
		log.Println(ksecDetailArray)
		for _, KsecDetail := range ksecDetailArray {
			queryUpdate := query.ProgramMapPloWithKsecQueryEntity{
				ID: KsecDetail.MapPLOID,
			}
			update := query.ProgramMapPloWithKsecQueryEntity{
				PloID:  update.ID,
				KsecID: KsecDetail.ID,
			}
			log.Println(KsecDetail.ID, KsecDetail.IsChecked)
			if pointer.GetBool(KsecDetail.IsChecked) {
				if KsecDetail.MapPLOID == nil {
					if err := tx.Create(&update).Error; err != nil {
						return err
					}
				}
			} else {
				if KsecDetail.MapPLOID != nil {
					if err := tx.Delete(&queryUpdate).Error; err != nil {
						return err
					}
				}
			}
		}

		if err := r.CreatePLODetail(tx, ploDetail.Children, ploId, ploDetail.ID, ksecDetails[index].Children); err != nil {
			return err
		}
	}
	return nil
}
