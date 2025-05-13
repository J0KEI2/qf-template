package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
)

func (rest *userDetailHandler) GetUserDetailPagination() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		options := helper.ExtractPaginationOption(c)

		respModel, err := rest.userDetailUseCase.GetUserDetailPagination(options)

		responseForm.Result = respModel

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
