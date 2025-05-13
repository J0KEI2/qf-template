package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetGeneralDetail(ProgramID uuid.UUID) (result *dto.ProgramGeneralDetailGetResponseDto, err error) {

	programMain := query.ProgramMainQueryEntity{
		ID: &ProgramID,
	}

	u.CommonRepository.GetFirst(&programMain, "ProgramGeneralDetail")

	programGeneralDetail := query.ProgramGeneralDetailQueryEntity{
		ID: programMain.ProgramGeneralDetailID,
	}

	if err = u.CommonRepository.GetFirst(&programGeneralDetail, "ProgramMajor", "ProgramMajor.ProgramPlanDetail", "ProgramMajor.ProgramPlanDetail.ProgramSubPlan"); err != nil {
		return
	}

	majorList := make([]dto.ProgramMajorDto, 0)
	for _, major := range programGeneralDetail.ProgramMajor {
		planDetailList := make([]dto.ProgramPlanDetailDto, 0)
		for _, plan := range major.ProgramPlanDetail {
			subPlanList := make([]dto.ProgramSubPlanDto, 0)
			for _, subPlan := range plan.ProgramSubPlan {
				subPlanList = append(subPlanList, dto.ProgramSubPlanDto{
					ID:                subPlan.ID,
					ProgramPlanDetail: subPlan.ProgramPlanDetailID,
					SubPlanName:       subPlan.SubPlanName,
					CreditRulesID:     subPlan.CreditRulesID,
					CreditRules:       subPlan.CreditRules,
					Credit:            subPlan.Credit,
				})
			}
			planDetailList = append(planDetailList, dto.ProgramPlanDetailDto{
				ID:                plan.ID,
				ProgramMajorID:    plan.ProgramMajorID,
				PlanName:          plan.PlanName,
				CreditRules:       plan.CreditRules,
				Credit:            plan.Credit,
				IsSplitPlan:       plan.IsSplitPlan,
				IsActive:          plan.IsActive,
				CreditRulesID:     plan.CreditRulesID,
				ProgramSubPlanDto: subPlanList,
			})
		}
		majorList = append(majorList, dto.ProgramMajorDto{
			ID:                major.ID,
			Name:              major.Name,
			ProgramPlanDetail: planDetailList,
		})
	}

	result = &dto.ProgramGeneralDetailGetResponseDto{
		ID:                    *programGeneralDetail.ID,
		UniversityName:        programGeneralDetail.UniversityName,
		FacultyID:             programGeneralDetail.FacultyID,
		ProgramNameTH:         programGeneralDetail.ProgramNameTH,
		ProgramNameEN:         programGeneralDetail.ProgramNameEN,
		ProgramCode:           programGeneralDetail.ProgramCode,
		DegreeNameTH:          programGeneralDetail.DegreeNameTH,
		DegreeNameEN:          programGeneralDetail.DegreeNameEN,
		DegreeNameShortenTH:   programGeneralDetail.DegreeNameShortenTH,
		DegreeNameShortenEN:   programGeneralDetail.DegreeNameShortenEN,
		OverallCredit:         programGeneralDetail.OverallCredit,
		ProgramMajorTypeID:    programGeneralDetail.ProgramMajorTypeID,
		ProgramMajorType:      programGeneralDetail.ProgramMajorType,
		ProgramDegreeTypeID:   programGeneralDetail.ProgramDegreeTypeID,
		ProgramDegreeType:     programGeneralDetail.ProgramDegreeType,
		NumberOfYear:          programGeneralDetail.NumberOfYear,
		ProgramLanguageID:     programGeneralDetail.ProgramLanguageID,
		ProgramLanguage:       programGeneralDetail.ProgramLanguage,
		Admission:             programGeneralDetail.Admission,
		MOU:                   programGeneralDetail.MOU,
		MOUFilepath:           programGeneralDetail.MOUFilepath,
		ProgramTypeID:         programGeneralDetail.ProgramTypeID,
		ProgramType:           programGeneralDetail.ProgramType,
		ProgramYearID:         programGeneralDetail.ProgramYearID,
		ProgramYear:           programGeneralDetail.ProgramYear,
		Semester:              programGeneralDetail.Semester,
		SemesterYear:          programGeneralDetail.SemesterYear,
		BoardApproval:         programGeneralDetail.BoardApproval,
		BoardApprovalDate:     programGeneralDetail.BoardApprovalDate,
		AcademicCouncil:       programGeneralDetail.AcademicCouncil,
		AcademicCouncilDate:   programGeneralDetail.AcademicCouncilDate,
		UniversityCouncil:     programGeneralDetail.UniversityCouncil,
		UniversityCouncilDate: programGeneralDetail.UniversityCouncilDate,
		IsSamePlanMajor:       programGeneralDetail.IsSamePlanMajor,
		ProgramMajor:          majorList,
		IsNationalProgram:     programGeneralDetail.IsNationalProgram,
		IsEnglishProgram:      programGeneralDetail.IsEnglishProgram,
		IsOther:               programGeneralDetail.IsOther,
		OtherName:             programGeneralDetail.OtherName,
		ProgramAdjustFrom:     programGeneralDetail.ProgramAdjustFrom,
		ProgramAdjustYear:     programGeneralDetail.ProgramAdjustYear,
	}

	return
}
