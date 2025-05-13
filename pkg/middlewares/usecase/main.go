package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type middlewaresUsecase struct {
	permissionUC domain.PermissionUseCase
	userUC       domain.UserUsecase
	roleUC       domain.RoleUseCase
	jwtResources *models.JwtResources
	commonRepository domain.CommonRepository
}

func NewMiddlewaresUsecase(permissionUC domain.PermissionUseCase, userUC domain.UserUsecase, roleUC domain.RoleUseCase, jwtResources *models.JwtResources, commonRepository domain.CommonRepository) domain.MiddlewaresUseCase {

	return &middlewaresUsecase{
		permissionUC,
		userUC,
		roleUC,
		jwtResources,
		commonRepository,
	}
}
