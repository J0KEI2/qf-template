package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) UpdateQF4Assessment(data *dto.QF4UpdateAssessmentRequestDto) (res *dto.QF4AssessmentResponseDto, err error) {

	updateData := query.QF4AssessmentQueryEntity{
		CategoryName:       &data.CategoryName,
		LearningAssessment: &data.LearningAssessment,
		Grade:              &data.Grade,
		GroupBased:         &data.GroupBased,
		Other:              &data.Other,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateAssessmentAction(&updateData, data.QF4ID))

	res = &dto.QF4AssessmentResponseDto{
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

func (u *qf4Usecase) updateAssessmentAction(dataUpdate *query.QF4AssessmentQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainQF4 := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}
		err = u.CommonRepository.GetFirst(&mainQF4)

		if err != nil {
			return err
		}

		if mainQF4.QF4AssessmentID == nil {
			return errors.New("can't find assessment for this qf")
		}

		queryUpdate := query.QF4Assessment{
			ID: mainQF4.QF4AssessmentID,
		}
		err = u.CommonRepository.Update(tx, queryUpdate, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
