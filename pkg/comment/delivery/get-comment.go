package delivery

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (h *commentHandler) GetComment() fiber.Handler {
	funcGetComment := func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		id := c.Params("id")
		idInt64, _ := strconv.ParseInt(id, 10, 64)
		response, err := h.CommentUseCase.GetComment(query.CommentQueryEntity{
			ID: uint(idInt64),
		})
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

	return funcGetComment
}
