package delivery

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h programHandler) CreateOrUpdateReference() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		programUID, err := uuid.Parse(c.Params("uuid", ""))

		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		userID, err := middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusUnauthorized,
				Title:   http.StatusText(http.StatusUnauthorized),
				Message: "Unauthorized",
			})
			return c.Status(http.StatusUnauthorized).JSON(responseForm)
		}

		request := dto.ProgramReferenceDto{}

		referenceIDStr := c.FormValue("id", "")
		if referenceIDStr != "" && referenceIDStr != "null" {
			referenceID, err := strconv.ParseUint(referenceIDStr, 10, 0)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid ref id")
			}

			refID := uint(referenceID)
			request.ID = &refID
		}

		referenceName := c.FormValue("reference_name", "")
		if referenceName != "" {
			request.ReferenceName = &referenceName
		}

		referenceDescription := c.FormValue("reference_description", "")
		if referenceDescription != "" {
			request.ReferenceDescription = &referenceDescription
		}

		// File upload example
		file, _ := c.FormFile("reference_file")

		referenceTypeIDStr := c.FormValue("reference_type_id", "null")
		if referenceTypeIDStr != "" && referenceTypeIDStr != "null" {
			referenceTypeID, err := strconv.Atoi(referenceTypeIDStr)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid reference_type_id ID")
			}

			request.ReferenceTypeID = &referenceTypeID
		}

		referenceTypeName := c.FormValue("reference_type_name", "null")
		if referenceTypeName != "" && referenceTypeName != "null" {
			request.ReferenceTypeName = &referenceTypeName
		}

		err = h.programUseCase.CreateOrUpdateReference(request, programUID, file, c, userID)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = map[string]interface{}{
			"resp_model": "Update success",
		}

		return c.JSON(responseForm)
	}
}
