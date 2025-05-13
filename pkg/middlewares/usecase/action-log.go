package usecase

import (
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zercle/kku-qf-services/internal/handlers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u *middlewaresUsecase) ActionLog() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		user_uid := new(string)
		tokenStr, err := handlers.ExtractBearerToken(c.Get(fiber.HeaderAuthorization))
		if err != nil {
			user_uid = nil
		} else {
			claims := new(dto.UserClaims)
			_, err := jwt.ParseWithClaims(tokenStr, claims, u.jwtResources.JwtKeyfunc)
			if err != nil {
				user_uid = nil
			}
			user_uid = &claims.UserID
		}

		apiPath, param, _ := strings.Cut(c.OriginalURL(), "?")
		statement := query.ActionLog{
			UserID:  user_uid,
			Method:  pointer.ToString(c.Route().Method),
			Action:  pointer.ToString(apiPath),
			Payload: pointer.ToString(string(c.Body())),
			Params:  pointer.ToString(param),
		}

		helper.ExecuteTransaction(u.commonRepository, u.createActionLogTx(&statement))
		return c.Next()
	}
}

func (u *middlewaresUsecase) createActionLogTx(statement *query.ActionLog) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.commonRepository.Create(tx, statement)
	}
}
