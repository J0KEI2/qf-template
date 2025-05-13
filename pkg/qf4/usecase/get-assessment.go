package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

)

func (u *qf4Usecase) GetQF4Assessment(qf4UidString string) (res *dto.QF4AssessmentResponseDto, err error) {
	qf4Uid, err := uuid.Parse(qf4UidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.QF4MainQueryEntity{
		QF4ID: &qf4Uid,
	}

	dest := query.QF4MainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "Assessment")

	res = &dto.QF4AssessmentResponseDto{
		ID:                 dest.Assessment.ID,
		CategoryName:       dest.Assessment.CategoryName,
		LearningAssessment: dest.Assessment.LearningAssessment,
		Grade:              dest.Assessment.Grade,
		GroupBased:         dest.Assessment.GroupBased,
		Other:              dest.Assessment.Other,
		CreatedAt:          dest.Assessment.CreatedAt,
		UpdatedAt:          dest.Assessment.UpdatedAt,
	}

	return
}
