package approvals

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h *programApprovalHandler) ApproveCurriculumCommitteeApprovalByProgramUID() fiber.Handler {
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

		request, updatedByUUID, err := prepareApproveApprovalRequest(c)
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

		approval, err := h.ApprovalUseCase.ApproveCurriculumCommitteeApprovalByProgramUID(&programUID, request, *userApprovalLevel)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		form, err := c.MultipartForm()
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		if err = h.ApprovalUseCase.UploadApprovalDocumentFile(c, programUID, form.File["files[]"], constant.APPROVAL_ATTRIBUTE_CURRICULUM_COMMITTEE, updatedByUUID, &approval.ID); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = approval

		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}
