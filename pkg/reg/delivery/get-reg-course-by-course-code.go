package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

func (h *regHandler) GetRegCourseByCourseCode() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		courseCode := c.Query("course_code")

		response, err := h.RegUseCase.GetRegCourseByCourseCode(courseCode)

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
}
