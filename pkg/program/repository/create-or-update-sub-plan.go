package repository

import (
	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateSubPlan(tx *gorm.DB, subPlans []dto.ProgramSubPlanDto, planDetailID *uint) (err error) {
	for _, subPlan := range subPlans {
		update := query.ProgramSubPlanQueryEntity{
			ProgramPlanDetailID: planDetailID,
			SubPlanName:         subPlan.SubPlanName,
			CreditRules:         subPlan.CreditRules,
			CreditRulesID:       subPlan.CreditRulesID,
			Credit:              subPlan.Credit,
		}
		if subPlan.ID != nil {
			queryPlan := query.ProgramSubPlanQueryEntity{
				ID:                  subPlan.ID,
				ProgramPlanDetailID: planDetailID,
				SubPlanName:         subPlan.SubPlanName,
				CreditRulesID:       subPlan.CreditRulesID,
				CreditRules:         subPlan.CreditRules,
				Credit:              subPlan.Credit,
			}
			if err = tx.Updates(&queryPlan).Error; err != nil {
				return err
			}
		} else {
			err = tx.Create(&update).Error
			if err != nil {
				return err
			}
			err = r.CreateSubPlanRelatedEmptyData(tx, pointer.GetUint(update.ID))
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (r programRepository) CreateSubPlanRelatedEmptyData(tx *gorm.DB, subPlanID uint) (err error) {
	ploFormat := query.ProgramPloFormatQueryEntity{
		ProgramSubPlanID: &subPlanID,
	}

	if err = tx.Create(&ploFormat).Error; err != nil {
		return
	}

	planAndEvaluate := query.ProgramPlanAndEvaluateQueryEntity{
		ProgramSubPlanID:      &subPlanID,
		StudentCharacteristic: pointer.ToString("[]"),
		ReceiveStudentPlan:    pointer.ToString("[]"),
		ProgramIncome:         pointer.ToString("[]"),
		ProgramOutcome:        pointer.ToString("[]"),
	}

	if err = tx.Create(&planAndEvaluate).Error; err != nil {
		return
	}

	return nil
}
