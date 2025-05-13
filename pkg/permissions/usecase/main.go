package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type PermissionUsecase struct {
	domain.PermissionRepository
	domain.CommonRepository
}

func NewRoleUsecase(repo domain.PermissionRepository, common domain.CommonRepository) domain.PermissionUseCase {
	return &PermissionUsecase{
		repo,
		common,
	}
}
