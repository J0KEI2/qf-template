package repository

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateYLOWithKsec(tx *gorm.DB, ploDetails []dto.ProgramYLOPLODto, yearAndSemesterID uint) error {
	for _, ploDetail := range ploDetails {
		for _, ksecDetail := range ploDetail.Ksec {
			err := r.createOrUpdateYLOWithKsec(tx, ksecDetail, yearAndSemesterID)
			if err != nil {
				return err
			}
		}
		for _, ploDetailChild := range ploDetail.Children {
			for _, ksecDetail := range ploDetailChild.Ksec {
				err := r.createOrUpdateYLOWithKsec(tx, ksecDetail, yearAndSemesterID)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (r programRepository) createOrUpdateYLOWithKsec(tx *gorm.DB, detail dto.YLOKsecDetail, yearAndSemesterID uint) error {
	if detail.MapYLOID != nil {
		if !*detail.IsChecked {
			yloQuery := query.ProgramYloWithKsecQueryEntity{
				ID: detail.MapYLOID,
			}

			err := tx.Delete(&yloQuery).Error
			if err != nil {
				return err
			}
		} else {
			statement := query.ProgramYloWithPloQueryEntity{
				ID: detail.MapYLOID,
			}

			yloKsecQuery := query.ProgramYloWithKsecQueryEntity{
				ProgramYearAndSemesterID: &yearAndSemesterID,
				ProgramMapPloWithKsecID:  detail.MapPLOID,
				Remark:                   detail.Remark,
				IsChecked:                detail.IsChecked,
			}
			err := tx.Where(statement).Updates(&yloKsecQuery).Error
			if err != nil {
				return err
			}
		}
	} else {
		if *detail.IsChecked {
			yloKsecQuery := query.ProgramYloWithKsecQueryEntity{
				ProgramYearAndSemesterID: &yearAndSemesterID,
				ProgramMapPloWithKsecID:  detail.MapPLOID,
				Remark:                   detail.Remark,
				IsChecked:                detail.IsChecked,
			}
			err := tx.Create(&yloKsecQuery).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
