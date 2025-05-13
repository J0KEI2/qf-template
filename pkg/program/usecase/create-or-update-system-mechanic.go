package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateSystemMechanic(systemAndMechanic dto.CreateOrUpdateSystemAndMechanicDto, programMainUID uuid.UUID) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateSystemMechanicTransaction(systemAndMechanic, programMainUID))
}

func (u programUsecase) CreateOrUpdateSystemMechanicTransaction(systemAndMechanic dto.CreateOrUpdateSystemAndMechanicDto, programMainUID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		programMain := query.ProgramMainQueryEntity{
			ID: &(programMainUID),
		}
		err = u.CommonRepository.GetFirst(&programMain)
		if err != nil {
			return err
		}

		queryPlo := query.ProgramSystemAndMechanicQueryEntity{
			ID: programMain.ProgramSystemAndMechanicID,
		}

		update := query.ProgramSystemAndMechanicQueryEntity{
			CourseExpectedAttribute: systemAndMechanic.CourseExpectedAttribute,
			CourseImprovingPlan:     systemAndMechanic.CourseImprovingPlan,
			CoursePolicies:          systemAndMechanic.CoursePolicies,
			CourseRisk:              systemAndMechanic.CourseRisk,
			CourseStrategies:        systemAndMechanic.CourseStrategies,
			CourseStudentComment:    systemAndMechanic.CourseStudentComment,
			MainContentAndStructure: systemAndMechanic.MainContentAndStructure,
		}
		return u.CommonRepository.Update(tx, queryPlo, &update)
	}
}
