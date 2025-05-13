package approvals

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h *programApprovalHandler) GetCurriculumCommitteeApprovalByProgramUID() fiber.Handler {
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

		userUID, err := middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusUnauthorized,
				Title:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			})
			return c.Status(http.StatusUnauthorized).JSON(responseForm)
		}

		userApprovalLevel, err := middlewares.GetProgramApprovalIDFromClaims(c)
		if err != nil || userApprovalLevel == nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusUnauthorized,
				Title:   http.StatusText(http.StatusUnauthorized),
				Message: err.Error(),
			})
			return c.Status(http.StatusUnauthorized).JSON(responseForm)
		}

		switch *userApprovalLevel {
		case uint(constant.CURRICULUM_COMMITEE_APPROVAL_LEVEL):
			approval, err := h.ApprovalUseCase.GetCurriculumCommitteeApprovalByProgramUID(&programUID, &userUID, *userApprovalLevel)
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

			// case uint(constant.CURRICULUM_COMMITEES_APPROVAL_LEVEL):
		default:
			approval, err := h.ApprovalUseCase.GetCurriculumCommitteesApprovalByProgramUID(&programUID, *userApprovalLevel)
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

			// default:
			// 	if *userApprovalLevel >= uint(constant.ADMIN_APPROVAL_LEVEL) {
			// 		approval, err := h.ApprovalUseCase.GetCurriculumCommitteesApprovalByProgramUID(&programUID, *userApprovalLevel)
			// 		if err != nil {
			// 			responseForm.Success = false
			// 			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
			// 				Code:    fiber.StatusInternalServerError,
			// 				Title:   "INTERNAL_SERVER_ERROR",
			// 				Message: err.Error(),
			// 			})
			// 			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
			// 		}

			// 		responseForm.Success = true
			// 		responseForm.Result = approval

			// 		return c.Status(fiber.StatusOK).JSON(responseForm)
			// 	} else {
			// 		responseForm.Success = false
			// 		responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
			// 			Code:    http.StatusUnauthorized,
			// 			Title:   http.StatusText(http.StatusUnauthorized),
			// 			Message: "Unathorized to get curriculum committee data.",
			// 		})
			// 		return c.Status(http.StatusUnauthorized).JSON(responseForm)
			// 	}
		}
	}
}
