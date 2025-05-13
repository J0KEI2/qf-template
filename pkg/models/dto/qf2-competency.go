package dto

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type ProgramCompetencyRequestDto struct {
	ProgramMainID         uuid.UUID                  `json:"program_main_id"`
	ProgramCompetencyList []ProgramCompetencyListDto `json:"items"`
}

type ProgramCompetencyListDto struct {
	ID                 *uint   `json:"id"`
	Order              *int    `json:"order"`
	SpecificCompetency *string `json:"specific_competency"`
	GenericCompetency  *string `json:"generic_competency"`
}

type ProgramCompetencyResponseDto struct {
	Items []ProgramCompetencyListDto `json:"items"`
	*models.PaginationOptions
}
