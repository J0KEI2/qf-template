package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdatePolicyAndStrategic(policyAndStrategic dto.ProgramPolicyAndStrategicRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdatePolicyAndStrategicTransaction(policyAndStrategic))
}

func (u programUsecase) CreateOrUpdatePolicyAndStrategicTransaction(policyAndStrategic dto.ProgramPolicyAndStrategicRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryTb := query.ProgramMainQueryEntity{
			ID: &policyAndStrategic.ProgramMainID,
		}

		u.CommonRepository.GetFirst(&queryTb)

		update := query.ProgramPolicyAndStrategicQueryEntity{
			ProgramPhilosophy: policyAndStrategic.ProgramPhilosophy,
			ProgramObjective:  policyAndStrategic.ProgramObjective,
			ProgramPolicy:     policyAndStrategic.ProgramPolicy,
			ProgramStrategic:  policyAndStrategic.ProgramStrategic,
			ProgramRisk:       policyAndStrategic.ProgramRisk,
			ProgramFeedback:   policyAndStrategic.ProgramFeedback,
		}

		if queryTb.ProgramPolicyAndStrategicID == nil {
			err = helper.ExecuteTransaction(u.CommonRepository, u.createPolicyAndStrategicAction(&update, policyAndStrategic.ProgramMainID))
			if err != nil {
				return err
			}
		} else {
			queryPolicyAndStrategic := query.ProgramPolicyAndStrategicQueryEntity{
				ID: queryTb.ProgramPolicyAndStrategicID,
			}

			err = u.CommonRepository.Update(tx, queryPolicyAndStrategic, &update)
			if err != nil {
				return err
			}
		}

		return
	}
}

func (u *programUsecase) createPolicyAndStrategicAction(data *query.ProgramPolicyAndStrategicQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		// update course assessment id in main table
		mainQuery := query.ProgramMainQueryEntity{
			ID: &mainUid,
		}

		mainUpdate := query.ProgramMainQueryEntity{
			ProgramPolicyAndStrategicID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
