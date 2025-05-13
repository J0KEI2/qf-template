package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (u *courseUsecase) GetCourseAssessment(courseUidString string) (res *dto.CourseAssessmentResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.CourseMainQueryEntity{
		CourseID: &courseUid,
	}

	dest := query.CourseMainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "Assessment")

	res = &dto.CourseAssessmentResponseDto{
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
