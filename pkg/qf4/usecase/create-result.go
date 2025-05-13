package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) CreateOrUpdateQF4Result(data *dto.QF4CreateResultRequestDto) (res *dto.QF4ResultResponseDto, err error) {
	queryTb := query.QF4MainQueryEntity{
		QF4ID: &data.QF4ID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	createOrUpdateData := query.QF4ResultQueryEntity{
		CategoryName:        &data.CategoryName,
		LearningOutcome:     &data.LearningOutcome,
		LearningExpectation: &data.LearningExpectation,
	}

	if queryTb.QF4ResultID == nil {
		err = helper.ExecuteTransaction(u.CommonRepository, u.createResultAction(&createOrUpdateData, data.QF4ID))
	} else {
		err = helper.ExecuteTransaction(u.CommonRepository, u.updateResultAction(&createOrUpdateData, data.QF4ID))
	}
	if err != nil {
		return nil, err
	}

	return &dto.QF4ResultResponseDto{
		ID:                  createOrUpdateData.ID,
		CategoryName:        createOrUpdateData.CategoryName,
		LearningOutcome:     createOrUpdateData.LearningOutcome,
		LearningExpectation: createOrUpdateData.LearningExpectation,
		CreatedAt:           createOrUpdateData.CreatedAt,
		UpdatedAt:           createOrUpdateData.UpdatedAt,
	}, nil
}

func (u *qf4Usecase) createResultAction(data *query.QF4ResultQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create course Result

		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		// update course assessment id in main table
		mainQuery := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}
		mainUpdate := query.QF4MainQueryEntity{
			QF4ResultID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
