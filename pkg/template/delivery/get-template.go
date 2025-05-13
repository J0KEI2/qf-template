package templateHandler

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

// TODO: don't forget to add call function in ./$MODULE_NAME/delivery/main.go
func (rest templateHandler) GetTemplate() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		demoID := c.Params("id")

		respModel, err := rest.templateUseCase.FetchTemplate(demoID)

		responseForm.Result = map[string]interface{}{
			"resp_model": respModel,
		}

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
