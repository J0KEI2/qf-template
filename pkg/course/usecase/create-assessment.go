package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
	"gorm.io/gorm"
)

func (u *courseUsecase) CreateOrUpdateCourseAssessment(data *dto.CourseCreateAssessmentRequestDto) (res *dto.CourseAssessmentResponseDto, err error) {
	queryTb := query.CourseMainQueryEntity{
		CourseID: &data.CourseID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	createOrUpdateData := query.CourseAssessmentQueryEntity{
		CategoryName:       &data.CategoryName,
		LearningAssessment: &data.LearningAssessment,
		Grade:              &data.Grade,
		GroupBased:         &data.GroupBased,
		Other:              &data.Other,
	}
	if queryTb.Assessment == nil {
		err = helper.ExecuteTransaction(u.CommonRepository, u.createAssessmentAction(&createOrUpdateData, data.CourseID))
	} else {
		err = helper.ExecuteTransaction(u.CommonRepository, u.updateAssessmentAction(&createOrUpdateData, data.CourseID))
	}

	if err != nil {
		return nil, err
	}

	return &dto.CourseAssessmentResponseDto{
		ID:                 createOrUpdateData.ID,
		CategoryName:       createOrUpdateData.CategoryName,
		LearningAssessment: createOrUpdateData.LearningAssessment,
		Grade:              createOrUpdateData.Grade,
		GroupBased:         createOrUpdateData.GroupBased,
		Other:              createOrUpdateData.Other,
		CreatedAt:          createOrUpdateData.CreatedAt,
		UpdatedAt:          createOrUpdateData.UpdatedAt,
	}, nil
}

func (u *courseUsecase) createAssessmentAction(data *query.CourseAssessmentQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create course assessment
		err = u.CommonRepository.Create(tx, data)

		if err != nil {
			return err
		}

		fmt.Println("\n mainUid: ", mainUid)
		// update course assessment id in main table
		mainQuery := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}
		mainUpdate := query.CourseMainQueryEntity{
			CourseAssessmentID: data.ID,
		}
		err = u.CommonRepository.Update(tx, &mainQuery, mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
