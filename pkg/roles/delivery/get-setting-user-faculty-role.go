package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
)

// TODO: don't forget to add call function in ./$MODULE_NAME/delivery/main.go
func (rest roleHandler) GetSettingUserFacultyRole() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		intRoleID := c.QueryInt("role_id", 0)
		intFacultyID := c.QueryInt("faculty_id", 0)
		if intRoleID == 0 {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: "Schema Invalid",
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		facultyID := uint(intFacultyID)
		roleID := uint(intRoleID)

		options := helper.ExtractPaginationOption(c)

		respModel, err := rest.roleUseCase.GetSettingUserFacultyRole(&roleID, &facultyID, options)

		responseForm.Result = respModel

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
