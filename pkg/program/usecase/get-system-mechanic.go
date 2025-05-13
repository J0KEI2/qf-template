package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetSystemMechanic(ProgramID uuid.UUID) (result *dto.ProgramSystemAndMechanicDto, err error) {
	programMain := query.ProgramMainQueryEntity{
		ID: &ProgramID,
	}

	err = u.CommonRepository.GetFirst(&programMain, "ProgramSystemAndMechanic")

	result = &dto.ProgramSystemAndMechanicDto{
		ID:                      programMain.ProgramSystemAndMechanic.ID,
		CoursePolicies:          programMain.ProgramSystemAndMechanic.CoursePolicies,
		CourseStrategies:        programMain.ProgramSystemAndMechanic.CourseStrategies,
		CourseRisk:              programMain.ProgramSystemAndMechanic.CourseRisk,
		CourseStudentComment:    programMain.ProgramSystemAndMechanic.CourseStudentComment,
		CourseExpectedAttribute: programMain.ProgramSystemAndMechanic.CourseExpectedAttribute,
		MainContentAndStructure: programMain.ProgramSystemAndMechanic.MainContentAndStructure,
		CourseImprovingPlan:     programMain.ProgramSystemAndMechanic.CourseImprovingPlan,
		CreatedAt:               programMain.ProgramSystemAndMechanic.CreatedAt,
		UpdatedAt:               programMain.ProgramSystemAndMechanic.UpdatedAt,
	}

	return
}
