package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func (h *userDetailHandler) SearchLecturerByName() fiber.Handler {
	funcSearchLecturerByName := func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		fullName := c.Query("fullName")

		criteria := models.LecturerFetchWithNameRequestModel{
			FullName: &fullName,
		}

		accessToken, err := h.hrUseCase.GetHRToken()

		if err != nil {
			responseForm.Success = false

			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error()})

			c.Status(http.StatusInternalServerError)

			return c.JSON(responseForm)
		}

		response, err := h.userDetailUseCase.SearchLecturerByName(*accessToken, &criteria)

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

	return funcSearchLecturerByName
}
