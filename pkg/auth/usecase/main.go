package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type authUsecase struct {
	domain.AuthRepository
}

func NewAuthUsecase(auth domain.AuthRepository) domain.AuthUseCase {
	return &authUsecase{
		auth,
	}
}
