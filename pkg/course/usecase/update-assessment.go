package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) UpdateCourseAssessment(data *dto.CourseUpdateAssessmentRequestDto) (res *dto.CourseAssessmentResponseDto, err error) {

	updateData := query.CourseAssessmentQueryEntity{
		CategoryName:       &data.CategoryName,
		LearningAssessment: &data.LearningAssessment,
		Grade:              &data.Grade,
		GroupBased:         &data.GroupBased,
		Other:              &data.Other,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateAssessmentAction(&updateData, data.CourseID))

	res = &dto.CourseAssessmentResponseDto{
		ID:                 updateData.ID,
		CategoryName:       updateData.CategoryName,
		LearningAssessment: updateData.LearningAssessment,
		Grade:              updateData.Grade,
		GroupBased:         updateData.GroupBased,
		Other:              updateData.Other,
		CreatedAt:          updateData.CreatedAt,
		UpdatedAt:          updateData.UpdatedAt,
	}

	return
}

func (u *courseUsecase) updateAssessmentAction(dataUpdate *query.CourseAssessmentQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainCourse := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}
		err = u.CommonRepository.GetFirst(&mainCourse)

		if err != nil {
			return err
		}

		if mainCourse.CourseAssessmentID == nil {
			return errors.New("can't find assessment for this qf")
		}

		queryUpdate := query.CourseAssessmentQueryEntity{
			ID: mainCourse.CourseAssessmentID,
		}
		err = u.CommonRepository.Update(tx, &queryUpdate, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
