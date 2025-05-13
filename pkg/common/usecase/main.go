package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type commonUsecase struct {
	domain.CommonRepository
}

func NewCommonUsecase(repo domain.CommonRepository) domain.CommonUseCase {
	return &commonUsecase{
		CommonRepository: repo,
	}
}
