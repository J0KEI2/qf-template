package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
)

func (rest *hrHandler) GetLecturerPagination() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		options := helper.ExtractPaginationOption(c)
		options.DefaultLimit(10)
		options.DefaultPage(1)

		respModel, err := rest.HRUseCase.GetLecturerPagination(options)
		if err != nil {
			responseForm.Success = false

			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error()})

			c.Status(http.StatusInternalServerError)

			return c.JSON(responseForm)
		}
		responseForm.Result = respModel

		responseForm.Success = true
		return c.JSON(responseForm)
	}
}
