package approvals

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u *programApprovalUsecase) CreateCHECOStatus(request dto.CreateCHECOStatusRequestDto, uid uuid.UUID) (checoId *uint, err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	queryUser := query.UserQueryEntity{
		UID: &uid,
	}

	if err = u.CommonRepository.GetFirst(&queryUser); err != nil {
		return nil, err
	}

	queryDb := query.CHECOQueryEntity{
		ProgramUID:   request.ProgramUID,
		NameEN:       queryUser.NameEN,
		NameTH:       queryUser.NameTH,
		StatusID:     request.StatusID,
		ApprovedDate: request.ApprovedDate,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, prepareCHECOCreateStatement(u, &queryDb)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return queryDb.ID, nil
}

func prepareCHECOCreateStatement(useCase *programApprovalUsecase, queryDb *query.CHECOQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return useCase.CommonRepository.Create(tx, queryDb)
	}
}
