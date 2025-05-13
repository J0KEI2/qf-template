package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type regHandler struct {
	domain.RegUseCase
}

func NewRegHandler(regRoute fiber.Router, regUseCase domain.RegUseCase) {

	handler := &regHandler{
		RegUseCase: regUseCase,
	}

	regRoute.Get("/token", handler.GetRegToken())
	regRoute.Get("/course", handler.GetRegCourseByCourseCode())
}
