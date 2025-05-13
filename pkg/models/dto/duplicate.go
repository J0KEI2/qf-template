package dto

import (
	"github.com/google/uuid"
)

type ProgramDuplicateRequestDto struct {
	ProgramMainID         uuid.UUID              `json:"program_main_id"`
}

type ProgramDuplicateResponseDto struct {
	ProgramMainID         uuid.UUID              `json:"program_main_id"`
}
