package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
)

func (h PermissionHandler) GetUserFacultyRolesByUid() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		userUID, err := uuid.Parse(c.Params("uid"))

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		options := helper.ExtractPaginationOption(c)

		respModel, err := h.permissionUsecase.GetFacultyRolesByUid(userUID, options)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Result = respModel
		responseForm.Success = true

		return c.JSON(responseForm)
	}
}
