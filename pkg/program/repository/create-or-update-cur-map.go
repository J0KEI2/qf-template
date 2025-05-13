package repository

import (
	"encoding/json"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"

	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateCurMapResp(tx *gorm.DB, paginationOptions *models.PaginationOptions, reqData dto.CreateOrUpdateCurMapRespRequestDto) (err error) {
	for _, fundamentalData := range reqData.ProgramFundamentalDetails.Items {
		for _, fundamentalPloData := range fundamentalData.ProgramCurMapRespDetails {
			fundamentalDataQuery := query.ProgramMapCurMapRespQueryEntity{
				ID:                fundamentalPloData.CurMapID,
				ProgramSubPlanID:      &reqData.ProgramSubPlanID,
				ProgramCourseDetailID: fundamentalData.ProgramCourseID,
				ProgramPloID:          fundamentalPloData.PLOID,
				Status:            fundamentalPloData.Status,
			}

			if err = tx.Updates(&fundamentalDataQuery).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					err = tx.Create(&fundamentalDataQuery).Error
				}
				if err != nil {
					return err
				}
			}
		}
	}

	for _, compulsoryData := range reqData.ProgramCompulsoryDetails.Items {
		for _, compulsoryPloData := range compulsoryData.ProgramCurMapRespDetails {
			compulsoryDataQuery := query.ProgramMapCurMapRespQueryEntity{
				ID:                compulsoryPloData.CurMapID,
				ProgramSubPlanID:      &reqData.ProgramSubPlanID,
				ProgramCourseDetailID: compulsoryData.ProgramCourseID,
				ProgramPloID:          compulsoryPloData.PLOID,
				Status:            compulsoryPloData.Status,
			}

			if err = tx.Updates(&compulsoryDataQuery).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					err = tx.Create(&compulsoryDataQuery).Error
				}
				if err != nil {
					return err
				}
			}
		}
	}

	for _, enrichmentData := range reqData.ProgramEnrichmentDetails.Items {
		for _, enrichmentPloData := range enrichmentData.ProgramCurMapRespDetails {
			enrichmentDataQuery := query.ProgramMapCurMapRespQueryEntity{
				ID:                enrichmentPloData.CurMapID,
				ProgramSubPlanID:      &reqData.ProgramSubPlanID,
				ProgramCourseDetailID: enrichmentData.ProgramCourseID,
				ProgramPloID:          enrichmentPloData.PLOID,
				Status:            enrichmentPloData.Status,
			}

			if err = tx.Updates(&enrichmentDataQuery).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					err = tx.Create(&enrichmentDataQuery).Error
				}
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r programRepository) CreateOrUpdateCurMapKsa(tx *gorm.DB, paginationOptions *models.PaginationOptions, reqData dto.CreateOrUpdateCurMapKsaRequestDto) (err error) {
	for _, fundamentalData := range reqData.ProgramFundamentalDetails.Items {
		for _, fundamentalKsaData := range fundamentalData.ProgramCurMapKsaDetails {
			ksaID, _ := json.Marshal(fundamentalKsaData.KsaID)
			fundamentalDataQuery := query.ProgramMapCurMapKsaQueryEntity{
				ID:                fundamentalKsaData.CurMapID,
				ProgramSubPlanID:      &reqData.ProgramSubPlanID,
				ProgramCourseDetailID: fundamentalData.ProgramCourseID,
				ProgramPloID:          fundamentalKsaData.PLOID,
				KsaID:             pointer.ToString(string(ksaID)),
			}

			if err = tx.Updates(&fundamentalDataQuery).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					err = tx.Create(&fundamentalDataQuery).Error
				}
				if err != nil {
					return err
				}
			}
		}
	}

	for _, compulsoryData := range reqData.ProgramCompulsoryDetails.Items {
		for _, compulsoryKsaData := range compulsoryData.ProgramCurMapKsaDetails {
			ksaID, _ := json.Marshal(compulsoryKsaData.KsaID)
			compulsoryDataQuery := query.ProgramMapCurMapKsaQueryEntity{
				ID:                compulsoryKsaData.CurMapID,
				ProgramSubPlanID:      &reqData.ProgramSubPlanID,
				ProgramCourseDetailID: compulsoryData.ProgramCourseID,
				ProgramPloID:          compulsoryKsaData.PLOID,
				KsaID:             pointer.ToString(string(ksaID)),
			}

			if err = tx.Updates(&compulsoryDataQuery).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					err = tx.Create(&compulsoryDataQuery).Error
				}
				if err != nil {
					return err
				}
			}
		}
	}

	for _, enrichmentData := range reqData.ProgramEnrichmentDetails.Items {
		for _, enrichmentKsaData := range enrichmentData.ProgramCurMapKsaDetails {
			ksaID, _ := json.Marshal(enrichmentKsaData.KsaID)
			enrichmentDataQuery := query.ProgramMapCurMapKsaQueryEntity{
				ID:                enrichmentKsaData.CurMapID,
				ProgramSubPlanID:      &reqData.ProgramSubPlanID,
				ProgramCourseDetailID: enrichmentData.ProgramCourseID,
				ProgramPloID:          enrichmentKsaData.PLOID,
				KsaID:             pointer.ToString(string(ksaID)),
			}

			if err = tx.Updates(&enrichmentDataQuery).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					err = tx.Create(&enrichmentDataQuery).Error
				}
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
