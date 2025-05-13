package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u *PermissionUsecase) GetCoreRolesByUid(uid uuid.UUID, options models.PaginationOptions) (result *dto.GetMapCoreRolesResponseDto, err error) {
	coreStatement := query.MapCoreRolesQueryEntity{
		UserID: &uid,
	}

	coreQuery := make([]query.MapCoreRolesQueryEntity, 0)

	err = u.CommonRepository.GetList(coreStatement, &coreQuery, &options, "Role")

	if err != nil {
		return nil, err
	}

	coreResult := make([]dto.MapCoreRoles, 0)

	for _, core := range coreQuery {
		coreResult = append(coreResult, dto.MapCoreRoles{
			ID:         core.ID,
			UserID:     core.UserID,
			RoleID:     core.RoleID,
			RoleNameTH: core.Role.RoleNameTH,
			RoleNameEN: core.Role.RoleNameEN,
		})
	}

	result = &dto.GetMapCoreRolesResponseDto{
		Items:             coreResult,
		PaginationOptions: &options,
	}
	return result, nil
}
