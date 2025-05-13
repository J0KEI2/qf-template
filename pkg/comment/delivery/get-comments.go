package delivery

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (h *commentHandler) GetComments() fiber.Handler {
	funcGetComment := func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		qfType := c.Params("qf_type")
		categoryType := c.Params("category_type")
		qfUID := c.Params("qf_uid")
		qfUUID, err := uuid.Parse(qfUID)
		showResolved := c.QueryBool("show_resolved", false)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   "BAD_REQUEST",
				Message: err.Error()})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		if qfType == "" || categoryType == "" {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   "BAD_REQUEST",
				Message: "invalid request"})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		var queryEnt query.CommentQueryEntity
		if showResolved {
			queryEnt = query.CommentQueryEntity{
				QFType:       qfType,
				QFUID:        qfUUID,
				CategoryType: categoryType,
				Resolve:      showResolved,
			}
		} else {
			queryEnt = query.CommentQueryEntity{
				QFType:       qfType,
				QFUID:        qfUUID,
				CategoryType: categoryType,
			}
		}

		fmt.Printf("\n >>>>>> query: %+v\n\n", queryEnt)
		response, err := h.CommentUseCase.GetComments(queryEnt)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error()})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = response

		return c.JSON(responseForm)
	}

	return funcGetComment
}
