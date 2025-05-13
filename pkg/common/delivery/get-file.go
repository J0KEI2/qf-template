package delivery

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

func (h commonHandler) GetFileByID() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		fileID, err := c.ParamsInt("file_id")
		if err != nil || fileID == 0 {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		filePath, err := h.commonUseCase.GetFilePathByID(uint(fileID))
		if err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		if err = c.SendFile(*filePath); err != nil {
			responseForm.Success = false
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusInternalServerError,
				Title:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		return

		// responseForm.Result = respModel

		// if err == nil {
		// 	responseForm.Success = true
		// }
		// return c.JSON(responseForm)
	}
}
