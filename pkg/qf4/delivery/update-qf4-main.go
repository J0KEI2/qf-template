package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (h *qf4Handler) UpdateQF4Main() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		req := new(dto.QF4UpdateMainRequestDto)

		if err := c.BodyParser(req); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		courseMain, err := h.qf4Usecase.UpdateQF4Main(req)

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = courseMain

		return c.JSON(responseForm)
	}
}
