package delivery

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (h programHandler) CreateOrUpdateCurriculumMapping() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		curMapType := c.Params("curMapType", "")

		paginationOptions := helper.ExtractPaginationOption(c)

		if strings.ToLower(curMapType) == "resp" {
			requestResp := dto.CreateOrUpdateCurMapRespRequestDto{}
			if err := c.BodyParser(&requestResp); err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    http.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: err.Error(),
				})
				return c.Status(http.StatusBadRequest).JSON(responseForm)
			}

			err = h.programUseCase.CreateOrUpdateCurriculumMappingResp(&paginationOptions, requestResp)
		} else if strings.ToLower(curMapType) == "ksa" {
			requestKsa := dto.CreateOrUpdateCurMapKsaRequestDto{}
			if err := c.BodyParser(&requestKsa); err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    http.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: err.Error(),
				})
				return c.Status(http.StatusBadRequest).JSON(responseForm)
			}

			err = h.programUseCase.CreateOrUpdateCurriculumMappingKsa(&paginationOptions, requestKsa)
		} else {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: "schema invalid",
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = map[string]interface{}{
			"resp_model": "Update success",
		}

		return c.JSON(responseForm)
	}
}
