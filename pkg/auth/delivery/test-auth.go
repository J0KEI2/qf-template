package delivery

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	migrate_models "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

type resposeBody struct {
	Token *jwt.Token           `json:"token"`
	User  migrate_models.Users `json:"user"`
}

func (h *authHandler) TestAuth() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.OauthResponse{}

		body := resposeBody{}
		fmt.Printf("\n >>>>>>>>>>>> authorization: %+v \n", c.Get(fiber.HeaderAuthorization))
		authHeader := strings.TrimSpace(c.Get(fiber.HeaderAuthorization))
		authHeaders := strings.Split(authHeader, " ")
		fmt.Printf("\n >>>>>>>>>>>> authHeaders: %+v \n", authHeaders)
		if len(authHeaders) != 2 || authHeaders[0] != "Bearer" {
			err = fiber.NewError(http.StatusUnauthorized, "Authorization: Bearer token")
			return
		}

		token := authHeaders[1]
		claims := new(dto.UserClaims)

		jwtToken, err := jwt.ParseWithClaims(token, claims, h.jwtResources.JwtKeyfunc)
		if err != nil {
			responseForm.Error = helpers.AccessDenied
			responseForm.ErrorDesc = err.Error()
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		body.Token = jwtToken

		body.User = migrate_models.Users{
			UID: uuid.MustParse(claims.UserID),
		}

		return c.JSON(body)
	}
}
