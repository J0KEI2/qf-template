package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateCurriculumMappingResp(paginationOptions *models.PaginationOptions, request dto.CreateOrUpdateCurMapRespRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateCurriculumMappingTransaction(paginationOptions, request))
}

func (u programUsecase) CreateOrUpdateCurriculumMappingKsa(paginationOptions *models.PaginationOptions, request dto.CreateOrUpdateCurMapKsaRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateCurriculumMappingKsaTransaction(paginationOptions, request))
}

func (u programUsecase) CreateOrUpdateCurriculumMappingTransaction(paginationOptions *models.PaginationOptions, reqData dto.CreateOrUpdateCurMapRespRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.ProgramRepository.CreateOrUpdateCurMapResp(tx, paginationOptions, reqData)
	}
}

func (u programUsecase) CreateOrUpdateCurriculumMappingKsaTransaction(paginationOptions *models.PaginationOptions, reqData dto.CreateOrUpdateCurMapKsaRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.ProgramRepository.CreateOrUpdateCurMapKsa(tx, paginationOptions, reqData)
	}
}

func (u programUsecase) InitCurriculumMapping(paginationOptions models.PaginationOptions, subPlanID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		courseQuery := query.ProgramCourseDetailQueryEntity{
			ProgramSubPlanID: &subPlanID,
		}

		courseList := []query.ProgramCourseDetailQueryEntity{}
		err := u.CommonRepository.GetList(&courseQuery, &courseList, nil)
		if err != nil {
			return err
		}

		programPLO := query.ProgramPloFormatQueryEntity{
			ProgramSubPlanID: &subPlanID,
		}

		if err = u.CommonRepository.GetFirst(&programPLO); err != nil {
			return err
		}

		ploDetails, err := u.RecursivePloDetail(programPLO.ID, pointer.ToUint(0))
		if err != nil {
			return err
		}

		fundamentalList, compulsoryList, enrichmentList, _ := u.SeparateCourseType(courseList, ploDetails, subPlanID)

		for _, fundamentalItem := range fundamentalList {
			err := u.CommonRepository.Create(tx, &fundamentalItem)
			if err != nil {
				return err
			}
		}

		for _, compulsoryItem := range compulsoryList {
			err := u.CommonRepository.Create(tx, &compulsoryItem)
			if err != nil {
				return err
			}
		}

		for _, enrichmentItem := range enrichmentList {
			err := u.CommonRepository.Create(tx, &enrichmentItem)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (u programUsecase) InitCurriculumMappingKsa(paginationOptions models.PaginationOptions, subPlanID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		courseQuery := query.ProgramCourseDetailQueryEntity{
			ProgramSubPlanID: &subPlanID,
		}

		courseList := []query.ProgramCourseDetailQueryEntity{}
		err := u.CommonRepository.GetList(&courseQuery, &courseList, nil)
		if err != nil {
			return err
		}

		programPLO := query.ProgramPloFormatQueryEntity{
			ProgramSubPlanID: &subPlanID,
		}

		if err = u.CommonRepository.GetFirst(&programPLO); err != nil {
			return err
		}

		ploDetails, err := u.RecursivePloDetail(programPLO.ID, pointer.ToUint(0))
		if err != nil {
			return err
		}

		fundamentalList, compulsoryList, enrichmentList, _ := u.SeparateCourseTypeKsa(courseList, ploDetails, subPlanID)

		for _, fundamentalItem := range fundamentalList {
			err := u.CommonRepository.Create(tx, &fundamentalItem)
			if err != nil {
				return err
			}
		}

		for _, compulsoryItem := range compulsoryList {
			err := u.CommonRepository.Create(tx, &compulsoryItem)
			if err != nil {
				return err
			}
		}

		for _, enrichmentItem := range enrichmentList {
			err := u.CommonRepository.Create(tx, &enrichmentItem)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
