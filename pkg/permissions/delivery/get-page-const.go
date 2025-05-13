package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/permission"
)

func (h PermissionHandler) GetPageConst() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		responseForm.Result = constant.PAGE_LIST
		responseForm.Success = true

		return c.JSON(responseForm)
	}
}
