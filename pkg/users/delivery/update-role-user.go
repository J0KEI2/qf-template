package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
)

// TODO: don't forget to add call function in ./$MODULE_NAME/delivery/main.go
func (h *userHandler) UpdateRoleUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		userID := c.Params("id")
		patchRequest := new(models.UpdateUserRoleRequest)

		if err := c.BodyParser(patchRequest); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
		}

		err = h.UserUsecase.UpdateRoleUser(userID, patchRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		responseForm.Result = map[string]interface{}{
			"message": "Update Lecturers Completed",
		}

		responseForm.Success = true

		return c.JSON(responseForm)
	}
}
