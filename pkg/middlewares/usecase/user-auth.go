package usecase

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/internal/handlers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *middlewaresUsecase) ExtractUserJwt() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		tokenStr, err := handlers.ExtractBearerToken(c.Get(fiber.HeaderAuthorization))
		if err != nil {
			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), err.Error())
		}

		claims := new(dto.UserClaims)
		jwtToken, err := jwt.ParseWithClaims(tokenStr, claims, u.jwtResources.JwtKeyfunc)
		if err != nil {
			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), err.Error())
		}
		if jwtToken != nil && jwtToken.Valid {
			c.Locals("userClaims", claims)
			c.Locals("token", jwtToken)
		} else {
			// debug
			log.Printf("%+v\nvalue: %+v", helpers.WhereAmI(), claims)
			log.Printf("%+v\nvalue: %+v", helpers.WhereAmI(), jwtToken)
			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), http.StatusText(http.StatusUnauthorized))
		}
		return c.Next()
	}
}
