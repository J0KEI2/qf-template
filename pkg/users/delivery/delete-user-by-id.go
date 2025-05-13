package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

func (h *userHandler) DeleteUserById() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		userID := c.Params("id")
		err = h.UserUsecase.DeleteUser(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		responseForm.Success = true
		responseForm.Result = map[string]interface{}{
			"status":  "OK",
			"message": "Delete User Success",
		}

		return c.JSON(responseForm)
	}
}
