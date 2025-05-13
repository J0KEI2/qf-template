package delivery

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (h programHandler) CreateOrUpdateLearningSolution() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		subPlanID, err := strconv.Atoi(c.Params("PloID", ""))

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		request := []dto.CreateOrUpdateLearningSolutionRequestDto{}

		if err := c.BodyParser(&request); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		err = h.programUseCase.CreateOrUpdatLearningSolution(request, uint(subPlanID))

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
