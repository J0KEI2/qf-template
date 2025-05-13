package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetStructure(structureId uint) (result *dto.ProgramStructureResponseDto, err error) {

	structure := query.ProgramStructureDetailQueryEntity{
		ID: &structureId,
	}

	u.CommonRepository.GetFirst(&structure, "Children")

	children := make([]dto.ProgramStructureResponseDto, 0)
	for _, structure := range structure.Children {
		children = append(children, dto.ProgramStructureResponseDto{
			ID:               structure.ID,
			ProgramSubPlanID: structure.ProgramSubPlanID,
			Name:             structure.Name,
			Order:            structure.Order,
			ParentID:         structure.ParentID,
			Qualification:    structure.Qualification,
			StructureCredit:  structure.StructureCredit,
		})
	}

	result = &dto.ProgramStructureResponseDto{
		ID:               structure.ID,
		ProgramSubPlanID: structure.ProgramSubPlanID,
		Name:             structure.Name,
		Order:            structure.Order,
		ParentID:         structure.ParentID,
		Children:         children,
		Qualification:    structure.Qualification,
		StructureCredit:  structure.StructureCredit,
	}
	return result, nil
}
