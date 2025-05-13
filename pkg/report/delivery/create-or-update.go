package delivery

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h reportHandler) CreateOrUpdateReport() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		userUID, err := middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Title:   "INTERNAL_SERVER_ERROR",
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		programUID, err := uuid.Parse(c.FormValue("uuid", ""))
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		var reportID *uint
		reportIDStr := c.FormValue("report_id", "null")
		if reportIDStr != "" && reportIDStr != "null" {
			reportIDUint, err := strconv.ParseUint(reportIDStr, 10, 0)
			if err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    http.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: err.Error(),
				})
				return c.Status(http.StatusBadRequest).JSON(responseForm)
			}
			reportIDs := uint(reportIDUint)
			reportID = &reportIDs
		}

		var reportName *string
		nameStr := c.FormValue("name", "null")
		if nameStr != "" && nameStr != "null" {
			reportName = &nameStr
		}

		var description *string
		descriptionStr := c.FormValue("description", "null")
		if descriptionStr != "" && descriptionStr != "null" {
			description = &descriptionStr
		}

		file, _ := c.FormFile("file")
		request := dto.CreateOrUpdateReportRequestDto{
			ProgramUID:  &programUID,
			ReportID:    reportID,
			Name:        reportName,
			Description: description,
			File:        file,
		}

		err = h.reportUseCase.CreateOrUpdateReport(c, request, userUID)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = "OK"

		return c.JSON(responseForm)
	}
}
