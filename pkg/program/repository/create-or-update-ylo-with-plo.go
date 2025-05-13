package repository

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateYLOWithPLO(tx *gorm.DB, ploDetails []dto.ProgramYLOPLODto, yearAndSemesterID uint) error {
	for _, yloPloData := range ploDetails {
		if yloPloData.MapYLOID != nil {
			if !*yloPloData.IsChecked {
				yloQuery := query.ProgramYloWithPloQueryEntity{
					ID: yloPloData.MapYLOID,
				}

				err := tx.Delete(&yloQuery).Error
				if err != nil {
					return err
				}
			} else {
				statement := query.ProgramYloWithPloQueryEntity{
					ID: yloPloData.MapYLOID,
				}

				yloPloQuery := query.ProgramYloWithPloQueryEntity{
					ID:                       yloPloData.MapYLOID,
					ProgramYearAndSemesterID: &yearAndSemesterID,
					ProgramPloID:             yloPloData.ID,
					Remark:                   yloPloData.Remark,
					IsChecked:                yloPloData.IsChecked,
				}
				err := tx.Where(&statement).Updates(&yloPloQuery).Error
				if err != nil {
					return err
				}
			}
		} else {
			if *yloPloData.IsChecked {
				yloPloQuery := query.ProgramYloWithPloQueryEntity{
					ProgramYearAndSemesterID: &yearAndSemesterID,
					ProgramPloID:             yloPloData.ID,
					Remark:                   yloPloData.Remark,
					IsChecked:                yloPloData.IsChecked,
				}
				err := tx.Create(&yloPloQuery).Error
				if err != nil {
					return err
				}
			}
		}

		if err := r.CreateOrUpdateYLOWithPLO(tx, yloPloData.Children, yearAndSemesterID); err != nil {
			return err
		}
	}
	return nil
}
