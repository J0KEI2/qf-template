package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type userDetailUseCase struct {
	userDetailRepo domain.UserDetailRepository
}

func NewUserDetailUseCase(repo domain.UserDetailRepository) domain.UserDetailUseCase {
	return &userDetailUseCase{
		userDetailRepo: repo,
	}
}
