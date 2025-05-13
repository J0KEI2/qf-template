package delivery

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)


func (h *commonHandler) GetFacultyByID() fiber.Handler {
	funcGetUserByID := func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		id, err := strconv.Atoi(c.Params("id", "0"))
		if err != nil {
			responseForm.Success = false

			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   "BAD_REQUEST",
				Message: err.Error()})

			c.Status(http.StatusBadRequest)

			return c.JSON(responseForm)
		}

		response, err := h.commonUseCase.GetFacultyByID(uint(id))

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

	return funcGetUserByID
}

