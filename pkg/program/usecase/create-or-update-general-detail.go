package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateGeneralDetail(generalDetail dto.ProgramGeneralDetailRequestDto) (err error) {
	err = helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateGeneralDetailTransaction(generalDetail))
	if err != nil {
		return err
	}
	year := pointer.GetInt(generalDetail.NumberOfYear)
	err = helper.ExecuteTransaction(u.CommonRepository, u.CreateSubPlanRelatedAction(generalDetail.ProgramMainID, uint(year)))
	if err != nil {
		return err
	}
	return
}

func (u programUsecase) CreateOrUpdateGeneralDetailTransaction(generalDetail dto.ProgramGeneralDetailRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryTb := query.ProgramMainQueryEntity{
			ID: &generalDetail.ProgramMainID,
		}

		u.CommonRepository.GetFirst(&queryTb)

		update := query.ProgramGeneralDetailQueryEntity{
			UniversityName:        generalDetail.UniversityName,
			FacultyID:             generalDetail.FacultyID,
			ProgramNameTH:         generalDetail.ProgramNameTH,
			ProgramNameEN:         generalDetail.ProgramNameEN,
			ProgramCode:           generalDetail.ProgramCode,
			DegreeNameTH:          generalDetail.DegreeNameTH,
			DegreeNameEN:          generalDetail.DegreeNameEN,
			ProgramMajorTypeID:    generalDetail.ProgramMajorTypeID,
			ProgramMajorType:      generalDetail.ProgramMajorType,
			ProgramDegreeTypeID:   generalDetail.ProgramDegreeTypeID,
			ProgramDegreeType:     generalDetail.ProgramDegreeType,
			DegreeNameShortenTH:   generalDetail.DegreeNameShortenTH,
			DegreeNameShortenEN:   generalDetail.DegreeNameShortenEN,
			IsSamePlanMajor:       generalDetail.IsSamePlanMajor,
			BranchNameTH:          generalDetail.BranchNameTH,
			BranchNameEN:          generalDetail.BranchNameEN,
			OverallCredit:         generalDetail.OverallCredit,
			NumberOfYear:          generalDetail.NumberOfYear,
			ProgramLanguageID:     generalDetail.ProgramLanguageID,
			ProgramLanguage:       generalDetail.ProgramLanguage,
			Admission:             generalDetail.Admission,
			MOU:                   generalDetail.MOU,
			MOUFilepath:           generalDetail.MOUFilepath,
			ProgramTypeID:         generalDetail.ProgramTypeID,
			ProgramType:           generalDetail.ProgramType,
			ProgramYearID:         generalDetail.ProgramYearID,
			ProgramYear:           generalDetail.ProgramYear,
			Semester:              generalDetail.Semester,
			SemesterYear:          generalDetail.SemesterYear,
			BoardApproval:         generalDetail.BoardApproval,
			BoardApprovalDate:     generalDetail.BoardApprovalDate,
			AcademicCouncil:       generalDetail.AcademicCouncil,
			AcademicCouncilDate:   generalDetail.AcademicCouncilDate,
			UniversityCouncil:     generalDetail.UniversityCouncil,
			UniversityCouncilDate: generalDetail.UniversityCouncilDate,
			IsNationalProgram:     generalDetail.IsNationalProgram,
			IsEnglishProgram:      generalDetail.IsEnglishProgram,
			IsOther:               generalDetail.IsOther,
			OtherName:             generalDetail.OtherName,
			ProgramAdjustFrom:     generalDetail.ProgramAdjustFrom,
			ProgramAdjustYear:     generalDetail.ProgramAdjustYear,
		}

		if queryTb.ProgramGeneralDetailID != nil {
			queryGeneralDetail := query.ProgramGeneralDetailQueryEntity{
				ID: queryTb.ProgramGeneralDetailID,
			}

			err = u.CommonRepository.Update(tx, queryGeneralDetail, &update)
			if err != nil {
				return err
			}
		} else {
			err = helper.ExecuteTransaction(u.CommonRepository, u.createGeneralDetailAction(&update, generalDetail.ProgramMainID))
			if err != nil {
				return err
			}
		}

		id := update.ID
		if id == nil {
			id = queryTb.ProgramGeneralDetailID
		}

		err = u.ProgramRepository.CreateOrUpdateMajor(tx, generalDetail.ProgramMajor, id, nil)
		if err != nil {
			return err
		}
		return
	}
}

func (u *programUsecase) createGeneralDetailAction(data *query.ProgramGeneralDetailQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		mainQuery := query.ProgramMainQueryEntity{
			ID: &mainUid,
		}

		mainUpdate := query.ProgramMainQueryEntity{
			ProgramGeneralDetailID: data.ID,
		}

		err = u.CommonRepository.Update(tx, &mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}

func (u *programUsecase) CreateSubPlanRelatedAction(programMainUid uuid.UUID, year uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryTb := query.ProgramMainQueryEntity{
			ID: &programMainUid,
		}

		u.CommonRepository.GetFirst(&queryTb)

		majorStatement := query.ProgramMajorQueryEntity{
			ProgramGeneralDetailID: queryTb.ProgramGeneralDetailID,
		}

		majorList := make([]query.ProgramMajorQueryEntity, 0)

		u.CommonRepository.GetList(&majorStatement, &majorList, nil, "ProgramPlanDetail.ProgramSubPlan.YearAndSemester")

		for _, majorItem := range majorList {
			for _, planDetailItem := range majorItem.ProgramPlanDetail {
				for _, subPlanItem := range planDetailItem.ProgramSubPlan {
					if len(subPlanItem.YearAndSemester) > 0 {
						continue
					}
					err = u.CreateYearAndSemesterByEducationYear(*subPlanItem.ID, year)
					if err != nil {
						return err
					}
				}
			}
		}
		return
	}
}
