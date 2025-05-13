package usecase

import (
	"sort"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetStructureBySubPlan(subPlanID uint, paginationOptions *models.PaginationOptions) (result *dto.GetCourseStructureResponseDto, err error) {
	structureResult, err := u.recursiveStructure(subPlanID, pointer.ToUint(0), paginationOptions) // get nil parentID
	if err != nil {
		return nil, err
	}
	result = &dto.GetCourseStructureResponseDto{
		Items:             structureResult,
		PaginationOptions: paginationOptions,
	}
	return result, nil
}

func (u programUsecase) recursiveStructure(subPlanID uint, parentID *uint, paginationOptions *models.PaginationOptions) (structuresResult []dto.ProgramStructureResponseDto, err error) {
	queryDb := query.ProgramStructureDetailQueryEntity{
		ProgramSubPlanID: &subPlanID,
		ParentID:         parentID,
	}

	structures := []query.ProgramStructureDetailQueryEntity{}

	if err = u.CommonRepository.GetListWithNilSearch(&queryDb, &structures, paginationOptions); err != nil {
		return nil, err
	}

	for _, structureCourse := range structures {
		children, err := u.recursiveStructure(subPlanID, structureCourse.ID, paginationOptions)
		if err != nil {
			return nil, err
		}
		structuresResult = append(structuresResult, dto.ProgramStructureResponseDto{
			ID:               structureCourse.ID,
			ProgramSubPlanID: structureCourse.ProgramSubPlanID,
			Name:             structureCourse.Name,
			Order:            structureCourse.Order,
			ParentID:         structureCourse.ParentID,
			Children:         children,
			Qualification:    structureCourse.Qualification,
			StructureCredit:  structureCourse.StructureCredit,
			CreatedAt:        structureCourse.CreatedAt,
			UpdatedAt:        structureCourse.UpdatedAt,
			DeletedAt:        structureCourse.DeletedAt,
		})
	}
	sort.SliceStable(structuresResult, func(i, j int) bool {
		return *structuresResult[i].Order < *structuresResult[j].Order
	})
	return structuresResult, nil
}
