package usecase

import (
	"sort"

	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
)

func (u commonUsecase) GetAllReferenceOption() (result []dto.ReferenceOption, err error) {
	record := make([]query.ReferenceOptionQueryEntity, 0)

	err = u.CommonRepository.GetList(query.ReferenceOptionQueryEntity{}, &record, nil)
	if err != nil {
		return nil, err
	}

	result = make([]dto.ReferenceOption, 0)

	for _, reference := range record {
		result = append(result, dto.ReferenceOption{
			ID:   reference.ID,
			Name: reference.Name,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return *result[i].ID < *result[j].ID
	})

	return result, nil
}
