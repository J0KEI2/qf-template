package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
)

func (h courseHandler) GetCoursePagination() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		options := helper.ExtractPaginationOption(c)
		options.DefaultLimit(10)

		// Todo : Filter by permission
		// userUID, err := middlewares.GetUserUIDFromClaims(c)
		// if err != nil {
		// 	responseForm.Success = false
		// 	responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
		// 		Code:    http.StatusBadRequest,
		// 		Title:   http.StatusText(http.StatusBadRequest),
		// 		Message: err.Error(),
		// 	})
		// 	return c.Status(http.StatusBadRequest).JSON(responseForm)
		// }

		// roleID, err := middlewares.GetRoleIDFromClaims(c)
		// if err != nil {
		// 	responseForm.Success = false
		// 	responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
		// 		Code:    http.StatusBadRequest,
		// 		Title:   http.StatusText(http.StatusBadRequest),
		// 		Message: err.Error(),
		// 	})
		// 	return c.Status(http.StatusBadRequest).JSON(responseForm)
		// }

		respModel, err := h.courseUsecase.GetMainCoursePagination(options)

		responseForm.Result = respModel

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
