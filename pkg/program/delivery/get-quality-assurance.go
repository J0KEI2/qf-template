package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
)

func (h programHandler) GetQualityAssurance() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		programUID, err := uuid.Parse(c.Params("uuid", ""))

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		respModel, err := h.programUseCase.GetQualityAssurance(programUID)

		responseForm.Result = respModel

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
