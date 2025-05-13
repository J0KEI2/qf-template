package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

// TODO: don't forget to add call function in ./$MODULE_NAME/delivery/main.go
func (h *userDetailHandler) CronUpdateLecturers() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		accessToken, err := h.hrUseCase.GetHRToken()
		if err != nil || accessToken == nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		err = h.userDetailUseCase.CronUpdateLecturers(*accessToken)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}

		responseForm.Result = map[string]interface{}{
			"message": "Update Lecturers Completed",
		}

		responseForm.Success = true

		return c.JSON(responseForm)
	}
}
