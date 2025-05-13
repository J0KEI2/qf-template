package usecase

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
	"gorm.io/gorm"
)

func (uc *middlewaresUsecase) SystemAuth(requiredSystemLevel uint) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		roleID, err := middlewares.GetRoleIDFromClaims(c)
		if err != nil {
			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), http.StatusText(http.StatusUnauthorized))
		}

		roleInfo, err := uc.roleUC.GetRoleByID(*roleID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return helpers.NewError(http.StatusInternalServerError, helpers.WhereAmI(), http.StatusText(http.StatusInternalServerError))
			}
			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), http.StatusText(http.StatusUnauthorized))
		}

		if *roleInfo.SystemLevel < requiredSystemLevel {
			// debug
			fmt.Printf("\n >>>>>>>>> *roleInfo.SystemLevel: %+v \n", *roleInfo.SystemLevel)
			fmt.Printf("\n >>>>>>>>> requiredSystemLevel: %+v \n", requiredSystemLevel)
			return helpers.NewError(http.StatusForbidden, helpers.WhereAmI(), http.StatusText(http.StatusForbidden))
		}

		return c.Next()
	}
}
