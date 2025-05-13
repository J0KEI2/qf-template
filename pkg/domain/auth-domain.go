package domain

import (
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

type AuthUseCase interface {
	ExtractThaidData(code string) (resp entity.ThaidTokenPostResponseEntity, responseError *helpers.ResponseError)
	GenerateJwtToken(user *models.UserFetchModel) (string, error)
}

type AuthRepository interface {
	DbAuthSVCMigrator() (err error)
	GetThaidDataFromDopa(code string) (resp entity.ThaidTokenPostResponseEntity, responseError *helpers.ResponseError)
}
