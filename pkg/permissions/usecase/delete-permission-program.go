package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) DeleteProgramPermission(deleteIdList []int) (err error) {
	functionArray := make([]func(tx *gorm.DB) error, 0)

	for _, deleteId := range deleteIdList {
		selectedId := deleteId
		queryDB := query.PermissionProgramQueryEntity{
			ID: &selectedId,
		}
		functionArray = append(functionArray, deleteProgramPermissionStatement(u, &queryDB))
	}

	return helper.ExecuteTransaction(u.CommonRepository, functionArray...)
}

func deleteProgramPermissionStatement(usecase *PermissionUsecase, queryDb *query.PermissionProgramQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return usecase.CommonRepository.Delete(tx, queryDb)
	}
}
