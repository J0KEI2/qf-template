package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) CreateOrUpdateCourseResult(data *dto.CourseCreateResultRequestDto) (res *dto.CourseResultResponseDto, err error) {
	if err != nil {
		return nil, err
	}

	queryTb := query.CourseMainQueryEntity{
		CourseID: &data.CourseID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	createOrUpdateData := query.CourseResultQueryEntity{
		CategoryName:        &data.CategoryName,
		LearningOutcome:     &data.LearningOutcome,
		LearningExpectation: &data.LearningExpectation,
	}

	if queryTb.CourseResultID == nil {
		err = helper.ExecuteTransaction(u.CommonRepository, u.createResultAction(&createOrUpdateData, data.CourseID))
	} else {
		err = helper.ExecuteTransaction(u.CommonRepository, u.updateResultAction(&createOrUpdateData, data.CourseID))
	}

	if err != nil {
		return nil, err
	}

	res = &dto.CourseResultResponseDto{
		ID:                  createOrUpdateData.ID,
		CategoryName:        createOrUpdateData.CategoryName,
		LearningOutcome:     createOrUpdateData.LearningOutcome,
		LearningExpectation: createOrUpdateData.LearningExpectation,
		CreatedAt:           createOrUpdateData.CreatedAt,
		UpdatedAt:           createOrUpdateData.UpdatedAt,
	}

	return
}

func (u *courseUsecase) createResultAction(data *query.CourseResultQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create course Result

		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		// update course assessment id in main table
		mainQuery := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}

		mainUpdate := query.CourseMainQueryEntity{
			CourseResultID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
