package delivery

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (h *commentHandler) CreateComment() fiber.Handler {

	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		createComment := models.CommentCreateQuery{}
		if err := c.BodyParser(&createComment); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		fmt.Printf("\n >>>> createComment: %+v \n\n", createComment)

		comment, err := h.CommentUseCase.CreateComment(query.CommentCreateEntity{
			QFType:       createComment.QFType,
			QFUID:        createComment.QFUID,
			CategoryType: createComment.CategoryType,
			Attribute:    createComment.Attribute,
			Commentator:  createComment.Commentator,
			Comments:     createComment.Comments,
			UpdatedBy:    createComment.UpdatedBy,
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
