package templateHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type templateHandler struct {
	templateUseCase domain.TemplateUseCase
}

func NewTemplateHandler(templateRoute fiber.Router, templateUseCase domain.TemplateUseCase) {

	handler := &templateHandler{
		templateUseCase: templateUseCase,
	}

	templateRoute.Get("/:bookID?", handler.GetTemplate())
}
