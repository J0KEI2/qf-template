package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type hrHandler struct {
	domain.HRUseCase
}

func NewHrHandler(hrRoute fiber.Router, hrUseCase domain.HRUseCase) {

	handler := &hrHandler{
		HRUseCase: hrUseCase,
	}

	hrRoute.Get("/token", handler.GetHRToken())
	hrRoute.Get("/educations", handler.GetHrEducationByEmail())
	hrRoute.Get("/employees", handler.GetLecturerPagination())
}
