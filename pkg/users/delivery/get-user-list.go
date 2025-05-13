package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

type Users struct {
	Email            *string `json:"email"`
	Name             *string `json:"name"`
	FacultyNameTH    *string `json:"facultyNameTH"`
	RoleSystemNameEN *string `json:"roleSystemNameEN"`
	Status           *string `json:"status"`
	Limit            *int    `json:"limit"`
	Page             *int    `json:"page"`
	OrderBy          *string `json:"orderBy"`
	Direction        *string `json:"direction"`
}

func (h *userHandler) GetUserList() fiber.Handler {
	funcGetUserList := func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		request := new(Users)

		if err := c.QueryParser(request); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			c.Status(http.StatusBadRequest)
			return c.JSON(responseForm)
		}

		options := helper.ExtractPaginationOption(c)

		result, err := h.UserUsecase.GetUserList(entity.UserFetchQueryEntity{
			Email:            request.Email,
			Name:             request.Name,
			RoleSystemNameEN: request.RoleSystemNameEN,
			FacultyNameTH:    request.FacultyNameTH,
			Status:           request.Status,
			OrderBy:          request.OrderBy,
			Direction:        request.Direction,
			Limit:            request.Limit,
			Page:             request.Page,
		}, &options)

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

		responseForm.Result = result

		return c.JSON(responseForm)
	}

	return funcGetUserList
}
