package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type roleUsecase struct {
	domain.RoleRepository
	domain.CommonRepository
}

func NewRoleUsecase(repo domain.RoleRepository, commonRepo domain.CommonRepository) domain.RoleUseCase {
	return &roleUsecase{
		repo,
		commonRepo,
	}
}
