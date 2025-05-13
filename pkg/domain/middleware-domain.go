package domain

import "github.com/gofiber/fiber/v2"

type MiddlewaresUseCase interface {
	ActionLog() fiber.Handler
	ExtractUserJwt() fiber.Handler
	SystemAuth(requiredSystemLevel uint) fiber.Handler
	ProgramActionAuth(reqProgramActionLevel uint) fiber.Handler
	ProgramApprovalAuth(reqProgramApprovalLevel uint) fiber.Handler
}
