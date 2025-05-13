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
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (h *authHandler) CurrentToken() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.OauthResponse{}
		type resposeBody struct {
			Token *jwt.Token          `json:"token"`
			User  migrateModels.Users `json:"user"`
			// Role  models.TbRole `json:"role"`
		}

		body := resposeBody{}

		fmt.Printf("\n >>>>>>>>>>>> authorization: %+v \n", c.Get(fiber.HeaderAuthorization))
		authHeader := strings.TrimSpace(c.Get(fiber.HeaderAuthorization))
		authHeaders := strings.Split(authHeader, " ")
		fmt.Printf("\n >>>>>>>>>>>> authHeaders: %+v \n", authHeaders)
		if len(authHeaders) != 2 || authHeaders[0] != "Bearer" {
			err = fiber.NewError(http.StatusUnauthorized, "Authorization: Bearer token")
			return
		}

		// body.Token =
		token := authHeaders[1]
		claims := new(dto.UserClaims)

		jwtToken, err := jwt.ParseWithClaims(token, claims, h.jwtResources.JwtKeyfunc)
		if err != nil {
			responseForm.Error = helpers.AccessDenied
			responseForm.ErrorDesc = err.Error()
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
			// return fiber.NewError(http.StatusUnauthorized, err.Error())
		}

		// if token, ok := c.Locals("token").(*jwt.Token); ok {
		body.Token = jwtToken
		// }

		// if user, ok := c.Locals("user").(migrateModels.Users); ok {
		body.User = migrateModels.Users{
			UID: uuid.MustParse(claims.UserID),
		}
		// }

		// if role, ok := c.Locals("role").(models.TbRole); ok {
		// 	body.Role = role
		// }

		return c.JSON(body)
	}
}
