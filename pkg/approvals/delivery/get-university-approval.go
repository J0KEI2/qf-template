package approvals

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h *programApprovalHandler) GetUniversityApprovalByProgramUID() fiber.Handler {
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

		userApprovalLevel, err := middlewares.GetProgramApprovalIDFromClaims(c)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusUnauthorized,
				Title:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			})
			return c.Status(http.StatusUnauthorized).JSON(responseForm)
		}

		approval, err := h.ApprovalUseCase.GetUniversityApprovalByProgramUID(&programUID, *userApprovalLevel)
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
