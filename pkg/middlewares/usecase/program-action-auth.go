package usecase

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
	"gorm.io/gorm"
)

func (uc *middlewaresUsecase) ProgramActionAuth(reqProgramActionLevel uint) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		roleID, err := middlewares.GetRoleIDFromClaims(c)
		if err != nil {
			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), http.StatusText(http.StatusUnauthorized))
		}

		roleInfo, err := uc.roleUC.GetRoleByID(*roleID)
		if err != nil {
			return helpers.NewError(http.StatusInternalServerError, helpers.WhereAmI(), http.StatusText(http.StatusInternalServerError))
		}

		if *roleInfo.ProgramActionLevel < reqProgramActionLevel {
			return helpers.NewError(http.StatusForbidden, helpers.WhereAmI(), http.StatusText(http.StatusForbidden))
		}

		// validate user
		userID, err := middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			return helpers.NewError(http.StatusInternalServerError, helpers.WhereAmI(), http.StatusText(http.StatusInternalServerError))
		}

		// ok, err := uc.roleUC.GetIfUserHasRoles(userID)
		ok, err := uc.roleUC.CheckIfUserHasRole(userID, roleID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return helpers.NewError(http.StatusInternalServerError, helpers.WhereAmI(), http.StatusText(http.StatusInternalServerError))
			}
			return helpers.NewError(http.StatusInternalServerError, helpers.WhereAmI(), http.StatusText(http.StatusInternalServerError))
		}

		if ok {
			return c.Next()
		} else {
			return helpers.NewError(http.StatusForbidden, helpers.WhereAmI(), http.StatusText(http.StatusForbidden))
		}
	}
}
