package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateCourseStructure(request dto.CreateOrUpdateCourseStructureRequestDto, subPlanID int) (err error) {
	subPlanIDUint := uint(subPlanID)
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateCourseStructureTransaction(request.Items, &subPlanIDUint))
}

func (u programUsecase) CreateOrUpdateCourseStructureTransaction(structures []dto.ProgramStructureRequestDto, subPlanID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.ProgramRepository.CreateOrUpdateCourseStructure(tx, structures, subPlanID, nil)
	}
}
