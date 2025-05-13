package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"
)

func (u *qf4Usecase) GetQF4Result(courseUidString string) (res *dto.QF4ResultResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.QF4MainQueryEntity{
		QF4ID: &courseUid,
	}

	dest := query.QF4MainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "Result")

	res = &dto.QF4ResultResponseDto{
		ID:                  dest.Result.ID,
		CategoryName:        dest.Result.CategoryName,
		LearningOutcome:     dest.Result.LearningOutcome,
		LearningExpectation: dest.Result.LearningExpectation,
		CreatedAt:           dest.Result.CreatedAt,
		UpdatedAt:           dest.Result.UpdatedAt,
	}

	return
}
