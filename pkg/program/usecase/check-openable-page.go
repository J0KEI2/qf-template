package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) CheckOpenablePage(ProgramID uuid.UUID) (result *dto.GetOpenablePageResponseDto, err error) {
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

	openableResponse := make([]dto.OpenablePageDto, 0)

	openableResponse = append(openableResponse,
		dto.OpenablePageDto{
			PageName:          "general_detail",
			Openable:          true,
			OpenableBySubPlan: nil,
		},
		dto.OpenablePageDto{
			PageName:          "policy_and_strategic",
			Openable:          true,
			OpenableBySubPlan: nil,
		},
	)

	subPlanIDsForProgramStructure := []int{}
	subPlanIDsForEducations := []int{}
	subPlanIDsForKsec := []int{}
	subPlanIDsForPlo := []int{}
	subPlanIDsForYlo := []int{}
	subPlanIDsForCmap := []int{}
	subPlanIDsForLearningProcess := []int{}
	subPlanIDsForPlanAndEvaluate := []int{}

	for _, major := range programGeneralDetail.ProgramMajor {
		for _, plan := range major.ProgramPlanDetail {
			if len(plan.ProgramSubPlan) > 0 {
				for _, subPlan := range plan.ProgramSubPlan {
					result, _ := u.GetPlo(*subPlan.ID)
					if result.Ksec != nil {
						subPlanIDsForPlo = append(subPlanIDsForPlo, int(*subPlan.ID))
					}

					if len(result.PLODetails) > 0 {
						subPlanIDsForYlo = append(subPlanIDsForYlo, int(*subPlan.ID))

					}

					yloResult, _ := u.GetYLODetail(*subPlan.ID)
					if len(yloResult.YLODetails) > 0 {
						subPlanIDsForCmap = append(subPlanIDsForCmap, int(*subPlan.ID))
					}

					subPlanIDsForProgramStructure = append(subPlanIDsForProgramStructure, int(*subPlan.ID))
					subPlanIDsForEducations = append(subPlanIDsForEducations, int(*subPlan.ID))
					subPlanIDsForKsec = append(subPlanIDsForKsec, int(*subPlan.ID))
					subPlanIDsForLearningProcess = append(subPlanIDsForLearningProcess, int(*subPlan.ID))
					subPlanIDsForPlanAndEvaluate = append(subPlanIDsForPlanAndEvaluate, int(*subPlan.ID))
				}
			}
		}
	}

	hasCmap := false
	hasYlo := false
	hasPlo := false
	if len(subPlanIDsForCmap) > 0 {
		hasCmap = true
	}

	if len(subPlanIDsForYlo) > 0 {
		hasYlo = true
	}

	if len(subPlanIDsForPlo) > 0 {
		hasPlo = true
	}

	openableResponse = append(openableResponse,
		dto.OpenablePageDto{
			PageName:          "program_structure",
			Openable:          true,
			OpenableBySubPlan: subPlanIDsForProgramStructure,
		},
		dto.OpenablePageDto{
			PageName:          "education_plan",
			Openable:          true,
			OpenableBySubPlan: subPlanIDsForEducations,
		},
		dto.OpenablePageDto{
			PageName:          "ksec",
			Openable:          true,
			OpenableBySubPlan: subPlanIDsForKsec,
		},
		dto.OpenablePageDto{
			PageName:          "plo",
			Openable:          hasPlo,
			OpenableBySubPlan: subPlanIDsForPlo,
		},
		dto.OpenablePageDto{
			PageName:          "ylo",
			Openable:          hasYlo,
			OpenableBySubPlan: subPlanIDsForYlo,
		},
		dto.OpenablePageDto{
			PageName:          "cmap",
			Openable:          hasCmap,
			OpenableBySubPlan: subPlanIDsForCmap,
		},
		dto.OpenablePageDto{
			PageName:          "learning_process",
			Openable:          true,
			OpenableBySubPlan: subPlanIDsForLearningProcess,
		},
		dto.OpenablePageDto{
			PageName:          "plan_and_evaluate",
			Openable:          true,
			OpenableBySubPlan: subPlanIDsForPlanAndEvaluate,
		},
		dto.OpenablePageDto{
			PageName:          "lecturer_owner",
			Openable:          true,
			OpenableBySubPlan: nil,
		},
		dto.OpenablePageDto{
			PageName:          "program_quality_assurance",
			Openable:          true,
			OpenableBySubPlan: nil,
		},
		dto.OpenablePageDto{
			PageName:          "program_mechanic",
			Openable:          true,
			OpenableBySubPlan: nil,
		},
		dto.OpenablePageDto{
			PageName:          "reference",
			Openable:          true,
			OpenableBySubPlan: nil,
		},
		dto.OpenablePageDto{
			PageName:          "document",
			Openable:          true,
			OpenableBySubPlan: subPlanIDsForPlanAndEvaluate,
		},
	)

	result = &dto.GetOpenablePageResponseDto{
		Items: openableResponse,
	}

	return
}
