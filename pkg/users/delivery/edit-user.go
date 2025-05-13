package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func (h *userHandler) EditUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		userID := c.Params("uid")
		patchRequest := new(models.PatchUserRequest)

		if err := c.BodyParser(patchRequest); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
		}

		if err := h.UserUsecase.EditUser(userID, *patchRequest); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		responseForm.Success = true

		return c.JSON(responseForm)
	}
}
