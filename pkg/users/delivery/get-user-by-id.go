package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
)

func (h *userHandler) GetUserByID() fiber.Handler {
	funcGetUserByID := func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		userID := c.Params("id")

		userUID, err := uuid.Parse(userID)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		response, err := h.UserUsecase.GetUserByID(userUID)

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error()})

			c.Status(http.StatusInternalServerError)

			return c.JSON(responseForm)
		}

		responseForm.Success = true

		responseForm.Result = response

		return c.JSON(responseForm)
	}

	return funcGetUserByID
}
