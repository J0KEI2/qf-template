package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (useCase qf4Usecase) DeleteQF4Main(data *dto.QF4DeleteMainRequestDto) (err error) {
	convertUid, _ := uuid.Parse(data.QF4UID)

	helper.ExecuteTransaction(useCase.CommonRepository, func(tx *gorm.DB) error {
		deleteQuery := query.QF4MainQueryEntity{
			QF4ID: &convertUid,
		}
		return useCase.CommonRepository.DeleteMainQFWithWhereClause(tx, &deleteQuery, "qf4_id", convertUid)
	})

	return nil
}
