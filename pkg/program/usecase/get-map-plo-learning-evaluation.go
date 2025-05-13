package usecase

import (
	"sort"

	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetMapPloWithLearningEvaluation(ploID uint, paginationOptions *models.PaginationOptions) (result dto.LearningEvaluationResponseDto, err error) {
	LearningEvaluationStatement := query.ProgramPLOLearningEvaluationQueryEntity{
		PloID: &ploID,
	}

	learningEvaluationQuery := make([]query.ProgramPLOLearningEvaluationQueryEntity, 0)

	if err = u.CommonRepository.GetList(LearningEvaluationStatement, &learningEvaluationQuery, paginationOptions); err != nil {
		return
	}

	learningEvaluations := make([]dto.LearningEvaluation, 0)

	for _, learningSolution := range learningEvaluationQuery {
		learningEvaluations = append(learningEvaluations, dto.LearningEvaluation{
			ID:        learningSolution.ID,
			PloID:     learningSolution.PloID,
			Order:     learningSolution.Order,
			Detail:    learningSolution.Detail,
			Key:       learningSolution.Key,
			CreatedAt: learningSolution.CreatedAt,
			UpdatedAt: learningSolution.UpdatedAt,
		})
	}

	sort.SliceStable(learningEvaluations, func(i, j int) bool {
		return *learningEvaluations[i].Order < *learningEvaluations[j].Order
	})

	result = dto.LearningEvaluationResponseDto{
		Items:             learningEvaluations,
		PaginationOptions: paginationOptions,
	}

	return
}
