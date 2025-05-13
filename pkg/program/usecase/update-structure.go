package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdateStructure(request dto.ProgramStructureRequestDto, structureId uint) (result *dto.ProgramStructureResponseDto, err error) {
	structureCredit, _ := request.StructureCredit.Int64()
	structureStatement := query.ProgramStructureDetailQueryEntity{
		ID: request.ID,
	}
	structureUpdate := query.ProgramStructureDetailQueryEntity{
		ProgramSubPlanID: request.SubPlanID,
		Name:             request.Name,
		Order:            request.Order,
		ParentID:         request.ParentID,
		Qualification:    request.Qualification,
		StructureCredit:  pointer.ToUint(uint(structureCredit)),
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdateStructureTransaction(&structureStatement, &structureUpdate))

	if err != nil {
		return nil, err
	}

	result = &dto.ProgramStructureResponseDto{
		ID:               structureUpdate.ID,
		ProgramSubPlanID: structureUpdate.ProgramSubPlanID,
		Name:             structureUpdate.Name,
		Order:            structureUpdate.Order,
		ParentID:         structureUpdate.ParentID,
		Qualification:    structureUpdate.Qualification,
		StructureCredit:  structureUpdate.StructureCredit,
		CreatedAt:        structureUpdate.CreatedAt,
		UpdatedAt:        structureUpdate.UpdatedAt,
		DeletedAt:        structureUpdate.DeletedAt,
	}

	return result, err
}

func (u programUsecase) UpdateStructureTransaction(statement *query.ProgramStructureDetailQueryEntity, update *query.ProgramStructureDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
