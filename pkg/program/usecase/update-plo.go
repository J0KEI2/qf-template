package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdatePlo(request dto.ProgramPLODetailDto, ploId uint) (result *dto.ProgramPLODetailDto, err error) {
	courseStatement := query.ProgramPloQueryEntity{
		ID: &ploId,
	}
	courseUpdate := query.ProgramPloQueryEntity{
		Order:     request.Order,
		ParentID:  request.ParentID,
		PLOPrefix: request.PLOPrefix,
		PLODetail: request.PLODetail,
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdatePloTransaction(&courseStatement, &courseUpdate))
	if err != nil {
		return nil, err
	}

	result = &dto.ProgramPLODetailDto{
		ID:        courseUpdate.ID,
		Order:     courseUpdate.Order,
		ParentID:  courseUpdate.ParentID,
		PLOPrefix: courseUpdate.PLOPrefix,
		PLODetail: courseUpdate.PLODetail,
		CreatedAt: courseUpdate.CreatedAt,
		UpdatedAt: courseUpdate.UpdatedAt,
		DeletedAt: courseUpdate.DeletedAt,
	}
	return
}

func (u programUsecase) UpdatePloTransaction(statement *query.ProgramPloQueryEntity, update *query.ProgramPloQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
