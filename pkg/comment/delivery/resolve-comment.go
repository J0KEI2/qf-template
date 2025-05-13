package delivery

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

func (h *commentHandler) ResolveComment() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		id := c.Params("id")
		idInt64, _ := strconv.ParseInt(id, 10, 64)

		comment, err := h.CommentUseCase.ResolveComment(uint(idInt64), true)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = comment

		return c.JSON(responseForm)
	}
}
