package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) CreateOrUpdateQF4Assessment(data *dto.QF4CreateAssessmentRequestDto) (res *dto.QF4AssessmentResponseDto, err error) {
	queryTb := query.QF4MainQueryEntity{
		QF4ID: &data.QF4ID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	createOrUpdateData := query.QF4AssessmentQueryEntity{
		CategoryName:       &data.CategoryName,
		LearningAssessment: &data.LearningAssessment,
		Grade:              &data.Grade,
		GroupBased:         &data.GroupBased,
		Other:              &data.Other,
	}
	if queryTb.Assessment == nil {
		err = helper.ExecuteTransaction(u.CommonRepository, u.createAssessmentAction(&createOrUpdateData, data.QF4ID))
	} else {
		err = helper.ExecuteTransaction(u.CommonRepository, u.updateAssessmentAction(&createOrUpdateData, data.QF4ID))
	}

	if err != nil {
		return nil, err
	}

	res = &dto.QF4AssessmentResponseDto{
		ID:                 createOrUpdateData.ID,
		CategoryName:       createOrUpdateData.CategoryName,
		LearningAssessment: createOrUpdateData.LearningAssessment,
		Grade:              createOrUpdateData.Grade,
		GroupBased:         createOrUpdateData.GroupBased,
		Other:              createOrUpdateData.Other,
		CreatedAt:          createOrUpdateData.CreatedAt,
		UpdatedAt:          createOrUpdateData.UpdatedAt,
	}

	return
}

func (u *qf4Usecase) createAssessmentAction(data *query.QF4AssessmentQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create qf4 assessment
		err = u.CommonRepository.Create(tx, data)

		if err != nil {
			return err
		}

		// update qf4 assessment id in main table
		mainQuery := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}
		mainUpdate := query.QF4MainQueryEntity{
			QF4AssessmentID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
