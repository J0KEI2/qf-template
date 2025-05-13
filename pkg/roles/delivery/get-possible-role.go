package delivery

import (
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

// TODO: don't forget to add call function in ./$MODULE_NAME/delivery/main.go
func (rest roleHandler) GetPossibleRole() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		userUID, err := middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			return err
		}

		respModel, err := rest.roleUseCase.GetPossibleRole(userUID)

		responseForm.Result = respModel

		if err == nil {
			responseForm.Success = true
		}
		return c.JSON(responseForm)
	}
}
