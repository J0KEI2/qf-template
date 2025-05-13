package usecase

import (
	"sort"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetCompetency(ProgramID uuid.UUID, paginationOptions *models.PaginationOptions) (result *dto.ProgramCompetencyResponseDto, err error) {

	programCompetency := query.ProgramCompetencyQueryEntity{
		ProgramMainID: &ProgramID,
	}

	competencyDest := []query.ProgramCompetencyQueryEntity{}
	u.CommonRepository.GetList(&programCompetency, &competencyDest, paginationOptions)

	competencyList := []dto.ProgramCompetencyListDto{}
	for _, competency := range competencyDest {
		competencyList = append(competencyList, dto.ProgramCompetencyListDto{
			ID:                 competency.ID,
			Order:              competency.Order,
			SpecificCompetency: competency.SpecificCompetency,
			GenericCompetency:  competency.GenericCompetency,
		})
	}

	sort.SliceStable(competencyList, func(i, j int) bool {
		return *competencyList[i].Order < *competencyList[j].Order
	})

	result = &dto.ProgramCompetencyResponseDto{
		Items:             competencyList,
		PaginationOptions: paginationOptions,
	}

	return
}
