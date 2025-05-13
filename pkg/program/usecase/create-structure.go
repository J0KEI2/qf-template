package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateStructure(request dto.ProgramStructureRequestDto, subPlanID uint) (result *dto.ProgramStructureResponseDto, err error) {
	structureChildren := make([]query.ProgramStructureDetailQueryEntity, 0)
	for _, children := range request.Children {
		structureCredit, _ := request.StructureCredit.Int64()
		structureChildren = append(structureChildren, query.ProgramStructureDetailQueryEntity{
			ProgramSubPlanID: &subPlanID,
			Name:             children.Name,
			Order:            children.Order,
			ParentID:         children.ParentID,
			Qualification:    children.Qualification,
			StructureCredit:  pointer.ToUint(uint(structureCredit)),
		})
	}
	structureCredit, _ := request.StructureCredit.Int64()
	structureQuery := query.ProgramStructureDetailQueryEntity{
		ProgramSubPlanID: &subPlanID,
		Name:             request.Name,
		Order:            request.Order,
		ParentID:         request.ParentID,
		Qualification:    request.Qualification,
		Children:         structureChildren,
		StructureCredit:  pointer.ToUint(uint(structureCredit)),
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.CreateStructureTransaction(&structureQuery, &subPlanID))

	if err != nil {
		return nil, err
	}

	structureChildrenResp := make([]dto.ProgramStructureResponseDto, 0)
	for _, children := range structureQuery.Children {
		structureChildrenResp = append(structureChildrenResp, dto.ProgramStructureResponseDto{
			ID:               children.ID,
			ProgramSubPlanID: children.ProgramSubPlanID,
			Name:             children.Name,
			Order:            children.Order,
			ParentID:         children.ParentID,
			Qualification:    children.Qualification,
			StructureCredit:  children.StructureCredit,
			CreatedAt:        children.CreatedAt,
			UpdatedAt:        children.UpdatedAt,
			DeletedAt:        children.DeletedAt,
		})
	}
	result = &dto.ProgramStructureResponseDto{
		ID:               structureQuery.ID,
		ProgramSubPlanID: structureQuery.ProgramSubPlanID,
		Name:             structureQuery.Name,
		Order:            structureQuery.Order,
		Children:         structureChildrenResp,
		ParentID:         structureQuery.ParentID,
		Qualification:    structureQuery.Qualification,
		StructureCredit:  structureQuery.StructureCredit,
		CreatedAt:        structureQuery.CreatedAt,
		UpdatedAt:        structureQuery.UpdatedAt,
		DeletedAt:        structureQuery.DeletedAt,
	}

	return result, err
}

func (u programUsecase) CreateStructureTransaction(structures *query.ProgramStructureDetailQueryEntity, subPlanID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Create(tx, structures)
	}
}
