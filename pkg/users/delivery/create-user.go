package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func (h *userHandler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		createUserDto := models.UserCreateQuery{}
		if err := c.BodyParser(&createUserDto); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		user, err := h.UserUsecase.CreateUser(createUserDto)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = user

		return c.JSON(responseForm)
	}
}
