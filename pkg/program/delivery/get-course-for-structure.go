package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
)

func (h programHandler) GetCourseForStructure() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		paginationOptions := helper.ExtractPaginationOption(c)

		respModel, err := h.programUseCase.GetCoursePaginationForStructure(paginationOptions)
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
