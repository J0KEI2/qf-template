package approvals

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h *programApprovalHandler) RejectFacultyApprovalByProgramUID() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		programUID, err := uuid.Parse(c.Params("program_uid", ""))
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		var request dto.RejectApprovalRequestDto
		if err := c.BodyParser(&request); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		// extract updatedBy from token if empty
		if request.UpdatedBy == "" {
			updatedBy, err := middlewares.GetUserUIDFromClaims(c)
			if err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    fiber.StatusInternalServerError,
					Title:   "INTERNAL_SERVER_ERROR",
					Message: err.Error(),
				})
				return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
			}
			request.UpdatedBy = updatedBy.String()
		}

		approval, err := h.ApprovalUseCase.RejectFacultyApprovalByProgramUID(&programUID, request)
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
		responseForm.Result = approval

		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}
