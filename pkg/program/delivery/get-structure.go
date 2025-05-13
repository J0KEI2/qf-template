package delivery

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

func (h programHandler) GetStructure() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		structureId, err := strconv.Atoi(c.Params("structureId"))

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		respModel, err := h.programUseCase.GetStructure(uint(structureId))
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Result = respModel
		responseForm.Success = true

		return c.JSON(responseForm)
	}
}
