package usecase

import (
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (u authUsecase) ExtractThaidData(code string) (resp entity.ThaidTokenPostResponseEntity, responseError *helpers.ResponseError) {
	return u.AuthRepository.GetThaidDataFromDopa(code)
}
