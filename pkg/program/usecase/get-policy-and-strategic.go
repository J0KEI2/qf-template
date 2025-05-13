package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetPolicyAndStrategic(ProgramID uuid.UUID) (result *dto.ProgramPolicyAndStrategicGetResponseDto, err error) {

	programMain := query.ProgramMainQueryEntity{
		ID: &ProgramID,
	}

	u.CommonRepository.GetFirst(&programMain, "ProgramPolicyAndStrategic")

	programPolicyAndStrategic := query.ProgramPolicyAndStrategicQueryEntity{
		ID: programMain.ProgramPolicyAndStrategicID,
	}

	u.CommonRepository.GetFirst(&programPolicyAndStrategic)

	result = &dto.ProgramPolicyAndStrategicGetResponseDto{
		ID:                *programPolicyAndStrategic.ID,
		ProgramMainID:         ProgramID,
		ProgramPhilosophy: programPolicyAndStrategic.ProgramPhilosophy,
		ProgramObjective:  programPolicyAndStrategic.ProgramObjective,
		ProgramPolicy:     programPolicyAndStrategic.ProgramPolicy,
		ProgramStrategic:  programPolicyAndStrategic.ProgramStrategic,
		ProgramRisk:       programPolicyAndStrategic.ProgramRisk,
		ProgramFeedback:   programPolicyAndStrategic.ProgramFeedback,
	}

	return
}
