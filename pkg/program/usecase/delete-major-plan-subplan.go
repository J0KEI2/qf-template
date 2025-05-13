package usecase

import (
	"log"

	"github.com/AlekSi/pointer"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) DeleteMajorPlanSubPlan(request dto.ProgramMajorAndPlanDeleteRequest) (err error) {

	complieTranactions := []func(tx *gorm.DB) error{}

	for _, majorId := range request.Majors {
		statement := query.ProgramMajorQueryEntity{
			ID: &majorId,
		}

		complieTranactions = append(complieTranactions, u.DeleteMajorTransaction(&statement))
	}

	for _, plan := range request.Plans {
		statement := query.ProgramPlanDetailQueryEntity{
			ID: &plan,
		}

		complieTranactions = append(complieTranactions, u.DeletePlanTransaction(&statement))
	}

	for _, subPlan := range request.Subplans {
		statement := query.ProgramSubPlanQueryEntity{
			ID: &subPlan,
		}

		if err = u.CommonRepository.GetFirst(&statement); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		}

		planStatement := query.ProgramPlanDetailQueryEntity{
			ID: statement.ProgramPlanDetailID,
		}

		if err = u.CommonRepository.GetFirst(&planStatement, "ProgramSubPlan"); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		}

		if len(planStatement.ProgramSubPlan) <= 1 {
			newSubplan := []dto.ProgramSubPlanDto{{
				SubPlanName:   planStatement.PlanName,
				CreditRulesID: planStatement.CreditRulesID,
				CreditRules:   planStatement.CreditRules,
				Credit:        planStatement.Credit,
			}}

			complieTranactions = append(complieTranactions, u.SplitPlanFromTrueToFalse(newSubplan, planStatement.ID))
		}

		complieTranactions = append(complieTranactions, u.DeleteSubPlanTransaction(&statement))
	}

	return helper.ExecuteTransaction(u.CommonRepository, complieTranactions...)
}

func (u programUsecase) DeleteMajorTransaction(query *query.ProgramMajorQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}

func (u programUsecase) DeletePlanTransaction(query *query.ProgramPlanDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}

func (u programUsecase) DeleteSubPlanTransaction(query *query.ProgramSubPlanQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Delete(tx, query)
	}
}

func (u programUsecase) SplitPlanFromTrueToFalse(subPlans []dto.ProgramSubPlanDto, planDetailID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		planStatement := query.ProgramPlanDetailQueryEntity{
			ID: planDetailID,
		}
		planUpdate := query.ProgramPlanDetailQueryEntity{
			IsSplitPlan: pointer.ToBool(false),
		}
		if err := u.CommonRepository.Update(tx, planStatement, &planUpdate); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		}
		return u.ProgramRepository.CreateOrUpdateSubPlan(tx, subPlans, planDetailID)
	}
}