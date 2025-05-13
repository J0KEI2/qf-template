package usecase

import (
	"sort"

	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetMapPloWithLearningSolution(ploID uint, paginationOptions *models.PaginationOptions) (result dto.LearningSolutionResponseDto, err error) {
	learningSolutionStatement := query.ProgramPLOLearningSolutionQueryEntity{
		PloID: &ploID,
	}

	learningSolutionQuery := make([]query.ProgramPLOLearningSolutionQueryEntity, 0)

	if err = u.CommonRepository.GetList(learningSolutionStatement, &learningSolutionQuery, paginationOptions); err != nil {
		return
	}

	learningSolutions := make([]dto.LearningSolution, 0)

	for _, learningSolution := range learningSolutionQuery {
		learningSolutions = append(learningSolutions, dto.LearningSolution{
			ID:        learningSolution.ID,
			PloID:     learningSolution.PloID,
			Order:     learningSolution.Order,
			Detail:    learningSolution.Detail,
			Key:       learningSolution.Key,
			CreatedAt: learningSolution.CreatedAt,
			UpdatedAt: learningSolution.UpdatedAt,
		})
	}

	sort.SliceStable(learningSolutions, func(i, j int) bool {
		return *learningSolutions[i].Order < *learningSolutions[j].Order
	})

	result = dto.LearningSolutionResponseDto{
		Items: learningSolutions,
		PaginationOptions: paginationOptions,
	}

	return
}
