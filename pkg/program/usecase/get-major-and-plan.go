package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetMajorAndPlan(ProgramID uuid.UUID) (result *dto.ProgramMajorAndPLanGetResponseDto, err error) {

	programMain := query.ProgramMainQueryEntity{
		ID: &ProgramID,
	}

	if err = u.CommonRepository.GetFirst(&programMain, "ProgramGeneralDetail"); err != nil {
		return
	}

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
				CreditRulesID:     plan.CreditRulesID,
				CreditRules:       plan.CreditRules,
				Credit:            plan.Credit,
				IsSplitPlan:       plan.IsSplitPlan,
				IsActive:          plan.IsActive,
				ProgramSubPlanDto: subPlanList,
			})
		}
		majorList = append(majorList, dto.ProgramMajorDto{
			ID:                major.ID,
			Name:              major.Name,
			ProgramPlanDetail: planDetailList,
		})
	}

	result = &dto.ProgramMajorAndPLanGetResponseDto{
		ProgramMajor: majorList,
	}

	return
}
