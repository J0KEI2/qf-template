package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) UpdateCourseResult(data *dto.CourseUpdateResultRequestDto) (res *dto.CourseResultResponseDto, err error) {

	updateData := query.CourseResultQueryEntity{
		CategoryName:        data.CategoryName,
		LearningOutcome:     data.LearningOutcome,
		LearningExpectation: data.LearningExpectation,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateResultAction(&updateData, *data.CourseID))

	res = &dto.CourseResultResponseDto{
		ID:                  updateData.ID,
		CategoryName:        updateData.CategoryName,
		LearningOutcome:     updateData.LearningOutcome,
		LearningExpectation: updateData.LearningExpectation,
		CreatedAt:           updateData.CreatedAt,
		UpdatedAt:           updateData.UpdatedAt,
	}

	return
}

func (u *courseUsecase) updateResultAction(dataUpdate *query.CourseResultQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainCourse := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}
		err = u.CommonRepository.GetFirst(&mainCourse)

		if err != nil {
			return err
		}

		if mainCourse.CourseResultID == nil {
			return errors.New("can't find assessment for this qf")
		}

		queryUpdate := query.CourseResult{
			ID: mainCourse.CourseResultID,
		}
		err = u.CommonRepository.Update(tx, queryUpdate, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
