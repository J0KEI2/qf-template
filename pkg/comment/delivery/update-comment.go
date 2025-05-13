package delivery

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (h *commentHandler) UpdateComment() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		id := c.Params("id")
		idInt64, _ := strconv.ParseInt(id, 10, 64)
		updateComment := models.CommentUpdateQuery{}
		if err := c.BodyParser(&updateComment); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		comment, err := h.CommentUseCase.UpdateComment(uint(idInt64),
			query.CommentUpdateEntity{
				QFType:       *updateComment.QFType,
				QFUID:        *updateComment.QFUID,
				CategoryType: *updateComment.CategoryType,
				Attribute:    *updateComment.Attribute,
				Commentator:  *updateComment.Commentator,
				Comments:     *updateComment.Comments,
				UpdatedBy:    *updateComment.UpdatedBy,
			})

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
