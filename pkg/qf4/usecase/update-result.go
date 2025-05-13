package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) UpdateQF4Result(data *dto.QF4UpdateResultRequestDto) (res *dto.QF4ResultResponseDto, err error) {

	updateData := query.QF4ResultQueryEntity{
		CategoryName:        data.CategoryName,
		LearningOutcome:     data.LearningOutcome,
		LearningExpectation: data.LearningExpectation,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateResultAction(&updateData, *data.QF4ID))

	res = &dto.QF4ResultResponseDto{
		ID:                  updateData.ID,
		CategoryName:        updateData.CategoryName,
		LearningOutcome:     updateData.LearningOutcome,
		LearningExpectation: updateData.LearningExpectation,
		CreatedAt:           updateData.CreatedAt,
		UpdatedAt:           updateData.UpdatedAt,
	}

	return
}

func (u *qf4Usecase) updateResultAction(dataUpdate *query.QF4ResultQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainQF4 := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}
		err = u.CommonRepository.GetFirst(&mainQF4)

		if err != nil {
			return err
		}

		if mainQF4.QF4ResultID == nil {
			return errors.New("can't find assessment for this qf")
		}

		queryUpdate := query.QF4Result{
			ID: mainQF4.QF4ResultID,
		}
		err = u.CommonRepository.Update(tx, queryUpdate, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
