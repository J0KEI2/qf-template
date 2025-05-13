package repository

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateMajor(tx *gorm.DB, majors []dto.ProgramMajorDto, generalDetailId *uint, parentID *uint) (err error) {
	for _, major := range majors {
		update := query.ProgramMajorQueryEntity{
			ProgramGeneralDetailID: generalDetailId,
			Name:               major.Name,
		}
		if major.ID != nil {
			queryMajor := query.ProgramMajorQueryEntity{
				ID:                 major.ID,
				ProgramGeneralDetailID: generalDetailId,
				Name:               major.Name,
			}
			if err = tx.Updates(&queryMajor).Error; err != nil {
				return err
			}

			if err = r.CreateOrUpdatePlanDetail(tx, major.ProgramPlanDetail, queryMajor.ID, nil); err != nil {
				return err
			}
		} else {
			err = tx.Create(&update).Error
			if err != nil {
				return err
			}
			if err = r.CreateOrUpdatePlanDetail(tx, major.ProgramPlanDetail, update.ID, nil); err != nil {
				return err
			}
		}
	}
	return nil
}
