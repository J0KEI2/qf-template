package repository

import (
	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdatePlanDetail(tx *gorm.DB, plans []dto.ProgramPlanDetailDto, majorId *uint, parentID *uint) (err error) {
	for _, plan := range plans {
		update := query.ProgramPlanDetailQueryEntity{
			ProgramMajorID: majorId,
			PlanName:       plan.PlanName,
			CreditRulesID:  plan.CreditRulesID,
			CreditRules:    plan.CreditRules,
			Credit:         plan.Credit,
			IsSplitPlan:    plan.IsSplitPlan,
			IsActive:       plan.IsActive,
		}

		subPlans := plan.ProgramSubPlanDto
		if !pointer.GetBool(plan.IsSplitPlan) || len(plan.ProgramSubPlanDto) <= 0 {
			subPlans = append(subPlans, dto.ProgramSubPlanDto{
				SubPlanName:   plan.PlanName,
				CreditRulesID: plan.CreditRulesID,
				CreditRules:   plan.CreditRules,
				Credit:        plan.Credit,
			})
		}

		if plan.ID != nil {
			queryPlan := query.ProgramPlanDetailQueryEntity{
				ID:             plan.ID,
				ProgramMajorID: majorId,
				PlanName:       plan.PlanName,
				CreditRulesID:  plan.CreditRulesID,
				CreditRules:    plan.CreditRules,
				Credit:         plan.Credit,
				IsSplitPlan:    plan.IsSplitPlan,
				IsActive:       plan.IsActive,
			}

			if err = tx.Updates(&queryPlan).Error; err != nil {
				return err
			}
			if err = r.CreateOrUpdateSubPlan(tx, subPlans, queryPlan.ID); err != nil {
				return err
			}

		} else {
			if err = tx.Create(&update).Error; err != nil {
				return err
			}
			if err = r.CreateOrUpdateSubPlan(tx, subPlans, update.ID); err != nil {
				return err
			}
		}

	}
	return nil
}
