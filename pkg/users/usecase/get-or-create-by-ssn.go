package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (u *userUsecase) GetOrCreateBySSN(userCreateQuery models.UserCreateQuery) (*entity.UserFetchEntity, error) {
	return u.userRepo.GetOrCreateBySSN(userCreateQuery)
}
