package approvals

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h *programApprovalHandler) CreateCHECOStatus() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {

		responseForm := helpers.ResponseForm{}
		request := dto.CreateCHECOStatusRequestDto{}

		programUID, err := uuid.Parse(c.FormValue("program_uid", ""))
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		statusID := c.FormValue("status_id", "")
		statusIDInt, err := strconv.Atoi(statusID)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: "stauts should be unsigned integer",
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}
		approvedDate := c.FormValue("approved_date", "")
		cleanedDateStr := strings.Split(approvedDate, " (")[0]
		approvedDateParsedTime, err := time.Parse(constant.FORM_DATA_TIME_FORMAT, cleanedDateStr)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		request.ProgramUID = &programUID
		request.StatusID = &statusIDInt
		request.ApprovedDate = &approvedDateParsedTime

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

		checoID, err := h.ApprovalUseCase.CreateCHECOStatus(request, userID)
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

		files := form.File["files[]"]
		err = h.ApprovalUseCase.UploadChecoDocumentFile(c, programUID, files, userID, checoID)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Success = true
		responseForm.Result = "OK"

		return c.Status(fiber.StatusCreated).JSON(responseForm)
	}
}
