package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (u *courseUsecase) GetCourseResult(courseUidString string) (res *dto.CourseResultResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.CourseMainQueryEntity{
		CourseID: &courseUid,
	}

	dest := query.CourseMainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "Result")

	res = &dto.CourseResultResponseDto{
		ID:                  dest.Result.ID,
		CategoryName:        dest.Result.CategoryName,
		LearningOutcome:     dest.Result.LearningOutcome,
		LearningExpectation: dest.Result.LearningExpectation,
		CreatedAt:           dest.Result.CreatedAt,
		UpdatedAt:           dest.Result.UpdatedAt,
	}

	return
}
