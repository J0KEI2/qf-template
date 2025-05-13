package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type userUsecase struct {
	userRepo   domain.UserRepository
	commonRepo domain.CommonRepository
}

func NewUserUsecase(userRepo domain.UserRepository, commonRepo domain.CommonRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		commonRepo: commonRepo,
	}
}
