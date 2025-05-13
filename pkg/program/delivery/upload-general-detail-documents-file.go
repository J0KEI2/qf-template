package delivery

import (
	"net/http"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h programHandler) UploadGeneralDetailMouDocumentsFile() fiber.Handler {
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

		var generalDetailId *uint
		generalDetailIdStr := c.FormValue("id", "null")
		if generalDetailIdStr != "" && generalDetailIdStr != "null" {
			generalDetailIdUint, err := strconv.ParseUint(generalDetailIdStr, 10, 0)
			if err != nil {
				responseForm.Success = false
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    http.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: err.Error(),
				})
				return c.Status(http.StatusBadRequest).JSON(responseForm)
			}

			generalId := uint(generalDetailIdUint)
			generalDetailId = &generalId
		}

		index := 0

		mouFileData := make([]dto.MouFileDto, 0)
		// Loop to get values like field[0], field[1], field[2]... until no value is found
		for {
			var fileID *uint
			fileIdKey := "mou_files[" + strconv.Itoa(index) + "][file_id]"
			fileIdVal := c.FormValue(fileIdKey, "")
			if fileIdVal == "" {
				break // No more values
			} else if fileIdVal != "null" {
				fileId, _ := strconv.ParseUint(fileIdVal, 10, 0)
				fileID = pointer.ToUint(uint(fileId))
			}
			// fmt.Println("\n fileIdVal: ", fileIdVal)

			fileKey := "mou_files[" + strconv.Itoa(index) + "][file]"
			file, err := c.FormFile(fileKey)
			// if not send file but only send file_id
			if err != nil && fileID != nil {
				mouFileData = append(mouFileData, dto.MouFileDto{
					FileID: fileID,
					File:   nil,
				})
			} else {
				mouFileData = append(mouFileData, dto.MouFileDto{
					FileID: fileID,
					File:   file,
				})
			}

			index++
		}
		_, err = h.programUseCase.UploadGeneralDetailMouDocumentsFile(c, programUID, generalDetailId, mouFileData, userUID)
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		// for _, fileID := range deleteFileIdList {
		// 	err = h.programUseCase.DeleteFileSystem(fileID)
		// 	if err != nil {
		// 		responseForm.Success = false
		// 		responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
		// 			Code:    http.StatusInternalServerError,
		// 			Title:   http.StatusText(http.StatusInternalServerError),
		// 			Message: err.Error(),
		// 		})
		// 		return c.Status(http.StatusInternalServerError).JSON(responseForm)
		// 	}
		// }

		responseForm.Result = "OK"
		responseForm.Success = true

		return c.JSON(responseForm)
	}
}
