package dto

import "github.com/google/uuid"

type ProgramPolicyAndStrategicRequestDto struct {
	ID                uint      `json:"id"`
	ProgramMainID         uuid.UUID `json:"program_main_id"`
	ProgramPhilosophy *string   `json:"program_philosophy"`
	ProgramObjective  *string   `json:"program_objective"`
	ProgramPolicy     *string   `json:"program_policy"`
	ProgramStrategic  *string   `json:"program_strategic"`
	ProgramRisk       *string   `json:"program_risk"`
	ProgramFeedback   *string   `json:"program_feedback"`
}

type ProgramPolicyAndStrategicGetResponseDto struct {
	ID                uint      `json:"id"`
	ProgramMainID         uuid.UUID `json:"program_main_id"`
	ProgramPhilosophy *string   `json:"program_philosophy"`
	ProgramObjective  *string   `json:"program_objective"`
	ProgramPolicy     *string   `json:"program_policy"`
	ProgramStrategic  *string   `json:"program_strategic"`
	ProgramRisk       *string   `json:"program_risk"`
	ProgramFeedback   *string   `json:"program_feedback"`
}
