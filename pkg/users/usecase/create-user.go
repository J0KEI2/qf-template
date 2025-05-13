package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (u *userUsecase) CreateUser(user models.UserCreateQuery) (userResult *entity.UserFetchEntity, err error) {
	return u.userRepo.CreateUser(user)
}