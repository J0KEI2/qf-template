package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) DeleteSystemPermission(uid uuid.UUID) (err error) {

	queryDB := query.PermissionSystemQueryEntity{
		UID: &uid,
	}

	helper.ExecuteTransaction(u.CommonRepository, deletePermissionStatement(u, &queryDB))

	return nil
}

func deletePermissionStatement(usecase *PermissionUsecase, queryDb *query.PermissionSystemQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return usecase.CommonRepository.Delete(tx, queryDb)
	}
}
