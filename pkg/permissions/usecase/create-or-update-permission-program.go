package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) CreateOrUpdateProgramPermission(permissionPrograms []dto.UpdatePermissionProgramDto) (result []dto.UpdatePermissionProgramDto, err error) {

	functionArray := make([]func(tx *gorm.DB) error, 0)

	for _, update := range permissionPrograms {
		updateProgramBody := query.PermissionProgramQueryEntity{
			UserUID:       update.UserUID,
			ProgramUID:    update.ProgramUID,
			Accessibility: update.Accessibility,
		}
		if update.ID != nil {
			updateProgramStatement := query.PermissionProgramQueryEntity{
				ID: update.ID,
			}
			functionArray = append(functionArray, updateProgramPermissionStatement(u, updateProgramStatement, &updateProgramBody))
		} else {
			functionArray = append(functionArray, createProgramPermissionStatement(u, &updateProgramBody))
		}
	}
	helper.ExecuteTransaction(u.CommonRepository, functionArray...)

	return permissionPrograms, nil
}

func updateProgramPermissionStatement(usecase *PermissionUsecase, queryDb query.PermissionProgramQueryEntity, updateBody *query.PermissionProgramQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return usecase.CommonRepository.Update(tx, queryDb, updateBody)
	}
}

func createProgramPermissionStatement(usecase *PermissionUsecase, createBody *query.PermissionProgramQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return usecase.CommonRepository.Create(tx, createBody)
	}
}
